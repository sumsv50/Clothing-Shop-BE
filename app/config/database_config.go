package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	cfg := mysql.Config{
		User:                 "u5pjo6qs39eqbiys",
		Passwd:               "yzTVUFLhdnhzxkKXpHYI",
		Net:                  "tcp",
		Addr:                 "bvczjxsdxdbwmd02n1cj-mysql.services.clever-cloud.com",
		DBName:               "bvczjxsdxdbwmd02n1cj",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	var DB *sql.DB
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connect to DB successfully!")
	return DB
}
