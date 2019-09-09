// Package datasource provides functionality to connect to data sources.
package datasource

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"

	"../logger"
)

// config represents yaml structure for datasources.yml configuration file
type config struct {
	RDBMS       []rdbmsConfig `yaml:"rdbms"`
	FileStorage []FileStorage `yaml:"storage-filesystem"`
	S3Storage   []S3Storage   `yaml:"storage-s3"`
}

// rdbmsConfig represent structure for data sources from rdbms database
type rdbmsConfig struct {
	Name            string `yaml:"name"`
	Driver          string `yaml:"driver"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	MaxIdleConn     *int   `yaml:"max_idle_connection"`
	MaxOpenConn     *int   `yaml:"max_open_connection"`
	MaxConnLifetime *int   `yaml:"max_connection_lifetime"`
}

// Logger
var log = logger.Get()

// Datasources
var datasources = make(map[string]interface{})

// NewRDBMS initiate connection to RDBMS database and returns rdbms datasource
func (cfg *rdbmsConfig) NewRDBMS() *sqlx.DB {
	// Check for optional values, set values if unset
	// If max idle connection is unset, set to 10
	if cfg.MaxIdleConn == nil {
		log.Debug("max_idle_connection is not set. Set to default: 10")
		cfg.MaxIdleConn = newInt(10)
	}
	// If max open connection is unset, set to 10
	if cfg.MaxOpenConn == nil {
		log.Debug("max_open_connection is not set. Set to default: 10")
		cfg.MaxOpenConn = newInt(10)
	}
	// If max idle connection is unset, set to 1 second
	if cfg.MaxConnLifetime == nil {
		log.Debug("max_connection_lifetime is not set. Set to default: 1 second")
		cfg.MaxConnLifetime = newInt(1)
	}
	// Create DSN string
	var dsn string
	switch cfg.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	default:
		fmt.Printf("Unsupported database driver. Driver: %s", cfg.Driver)
		os.Exit(13)
	}
	// Create connection
	log.Debug("Connecting to database")
	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database. Driver: %s. Error: %s", cfg.Driver, err.Error())
		os.Exit(14)
	}
	log.Debug("Connected to database")
	// Test ping
	log.Debug("Ping to database")
	if err := db.Ping(); err != nil {
		fmt.Printf("Unable to ping to database. Driver: %s. Error: %s", cfg.Driver, err.Error())
		os.Exit(15)
	}
	// Set lifetime
	maxLifetime := time.Duration(*cfg.MaxConnLifetime) * time.Second
	db.SetConnMaxLifetime(maxLifetime)
	// Set open connection
	db.SetMaxOpenConns(*cfg.MaxOpenConn)
	// Set idle connection
	db.SetMaxIdleConns(*cfg.MaxIdleConn)
	// Return database instance
	return db
}

// Init load datasources configuration from file and initiate datasource instances.
func Init() {
	// Load datasource definition in file
	var c config
	bytes, err := ioutil.ReadFile("datasources.yml")
	if err != nil {
		fmt.Printf("Unable to read datasources.yml file. Error: %s\n", err.Error())
		os.Exit(10)
	}
	// Parse error codes file
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		fmt.Printf("Unable to parse datasources.yml file. Error: %s\n", err.Error())
		os.Exit(11)
	}
	// Init data sources from database
	for _, dbConf := range c.RDBMS {
		log.Debugf("Init data sources from database. Name: %s", dbConf.Name)
		datasources[dbConf.Name] = dbConf.NewRDBMS()
	}
	// Init storage datasources
	for _, sfConf := range c.FileStorage {
		name := sfConf.Name
		log.Debugf("Init file storage datasources. Name: %s", name)
		datasources[name] = sfConf.Init()
	}
	// Init s3 storage
	for _, v := range c.S3Storage {
		name := v.Name
		log.Debugf("Init file s3 storage datasources. Name: %s", name)
		datasources[name] = v.Init()
	}
}

// GetRDBMS retrieve instance of RDBMS-based datasource
func GetRDBMS(name string) *sqlx.DB {
	return datasources[name].(*sqlx.DB)
}

func GetFileUploader(name string) FileUploader {
	return datasources[name].(FileUploader)
}

// newInt returns pointer of a variable that contains integer
func newInt(i int) *int {
	return &i
}
