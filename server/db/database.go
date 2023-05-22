package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlConnection() *sql.DB {
	//config based
	dbUrl := "root:Digi@2023@/infraguard_manager?parseTime=true"
	log.Print("Create MySql Connection")
	sql, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Println("Cannot connect to database", err)
	}
	return sql
}
