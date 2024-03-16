package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
	SqlDB *sql.DB
}

func Get() *DB {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	var err error
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database:", cfg.FormatDSN())
	return &DB{SqlDB: sqlDB}
}

func (db *DB) Clear() {
	db.SqlDB.Exec("call clear_db;")
}
