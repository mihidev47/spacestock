package repository

import (
	"../datasource"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var uploader datasource.FileUploader

var App AppRepository
var Apartment AptRepository

const (
	FolderAvatar        = "avatars"
	FolderChefCover     = "chef-covers"
	FolderCustomerCover = "customer-covers"
	FolderDocument      = "documents"
	FolderMenu          = "menus"
)

func Init() {
	// Get datasource
	db = datasource.GetRDBMS("koki_db")
	// Initiate repository
	App = &appRepository{db: db}
	Apartment = initAptRepository()
}

// BeginTx initiate database SQL transaction.
// BeginTx MUST be followed by deferred call of ReleaseTx
func BeginTx() (*sqlx.Tx, error) {
	return db.Beginx()
}
