package data

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func Initialize(dbPath string) error {
	if dbPath == ":memory:" {
		log.Println("WARN the current instance of this server is running without database, if the server, restart everything is lost. To set up a database path, use -db argument")
	}
	log.Printf("INFO loading database file at %s", dbPath)
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	return nil
}
