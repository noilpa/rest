package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

// Singleton db connection
var DB *sql.DB = nil

func InitDb() error {

	var err error
	uri := os.Getenv("DB_URI")
	DB, err = sql.Open("postgres", uri)
	if err != nil {
		return err
	}

	// verifies connection to the database
	err = DB.Ping()
	if err != nil {
		return err
	}


	return nil

}