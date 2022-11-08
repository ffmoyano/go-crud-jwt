package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var db *sql.DB

func Open() {
	var err error
	driver := os.Getenv("dbDriver")
	connString := os.Getenv("dburl")
	db, err = sql.Open(driver, connString)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(60 * time.Second)
	db.SetConnMaxLifetime(30 * time.Minute)
}

func Get() *sql.DB {
	return db
}
