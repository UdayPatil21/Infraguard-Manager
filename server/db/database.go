package db

import (
	"infraguard-manager/helpers/configHelper"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	ActivationDB = "AgentActivations"
	ServerDB     = "Servers"
)

//My SQL driver
// func MySqlConnection() *sql.DB {
// 	dburl := configHelper.GetString("DBURL")
// 	//read config based db url and connect
// 	log.Print("Create MySql Connection")
// 	sql, err := sql.Open("mysql", dburl)
// 	if err != nil {
// 		log.Println("Cannot connect to database", err)
// 	}
// 	return sql
// }
func MySqlConnection() *gorm.DB {

	dburl := configHelper.GetString("DBURL")
	//read config based db url and connect
	log.Print("Create MySql Connection")
	gorm, err := gorm.Open("mysql", dburl)
	if err != nil {
		log.Println("Cannot connect to database", err)
	}
	return gorm
}
