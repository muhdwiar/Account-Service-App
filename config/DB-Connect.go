package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(root string) *sql.DB {
	db, err := sql.Open("mysql", root)

	if err != nil {
		log.Fatal("Error Conecting :", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
		// panic("error connect db")
	} else {
		fmt.Println("success connect to DB")
	}

	return db
}
