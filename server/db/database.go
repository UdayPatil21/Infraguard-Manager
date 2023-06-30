package db

import (
	"infraguard-manager/helpers/configHelper"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	ActivationDB = "Clusters"
	ServerDB     = "Servers"
)

var DBInstance *gorm.DB
var DBError error

//Initialized mysql connection
func Init() {
	//Create one time database connection and share it to the application
	//Database connection pooling
	/*
		is a way to reduce the cost of opening and closing connections by maintaining a “pool” of open connections
		that can be passed from database operation to database operation as needed.
	*/
	MySqlConnection()
}

func MySqlConnection() {
	dburl := configHelper.GetString("DBURL")
	//read config based db url and connect
	DBInstance, DBError = gorm.Open("mysql", dburl)
	if DBError != nil {
		log.Println("Cannot connect to database", DBError)
	}
	log.Println("Connected to Database!")
}
