package db

import (
	"database/sql"
	"infraguard-manager/helpers/configHelper"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlConnection() *sql.DB {
	dburl := configHelper.GetString("DBURL")
	//read config based db url and connect
	log.Print("Create MySql Connection")
	sql, err := sql.Open("mysql", dburl)
	if err != nil {
		log.Println("Cannot connect to database", err)
	}
	return sql
}
