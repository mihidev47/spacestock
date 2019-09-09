package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"./config"
	"./crypto"
	"./datasource"
	"./flags"
	"./logger"
	"./repository"
	"./router"
	"./service"
	"./util"
)

var log = logger.Get()

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "Version")
	configFile := flag.String("c", "config.yml", "config.yml")
	// Parse flags
	flag.Parse()
	// If show version is true, print application version and exit
	if *showVersion {
		printVersion()
		os.Exit(0)
	}
	// If config file is unset, exit
	if *configFile == "" {
		fmt.Println("Configuration file is not specified. Use -c config.yml to specify one")
		os.Exit(1)
	}
	// Get starting time
	start := time.Now()
	// Load config file
	config.Init(*configFile)
	// Initiate logger
	logger.Init()
	// Initiate utilities
	util.Init()
	// Initiate datasources
	datasource.Init()
	// Initiate crypto
	crypto.Init()
	// Initiate repository
	repository.Init()
	// Initiate services
	service.Init()
	// Get router and start server
	r := router.New(start)
	// Start server
	log.Infof("Boot time: %s", time.Since(start))
	log.Fatal(http.ListenAndServe(getPort(), r))
}

// getPort retrieve port from config
func getPort() string {
	port := config.MustGetString("server.port")
	return ":" + port
}

func printVersion() {
	fmt.Printf("%s version %s, build %s\n", flags.AppName, flags.AppVersion, flags.AppCommitHash)
}
