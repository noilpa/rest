package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// hardcode for simplicity
// move it to OS envs
const (
	dbhost = "localhost"
	dbport = "5433"
	dbuser = "postgres"
	dbpass = "postgres"
	dbname = "restream"
)

// Singleton db connection
var DB *sql.DB = nil

func InitDb() error {

	connStr := fmt.Sprintf(
		"port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbport,
		dbuser,
		dbname,
		dbpass,
		)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// verifies connection to the database
	if err := DB.Ping(); err != nil {
		return err
	}


	return nil

}