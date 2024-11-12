package configs

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func GetConnection() *sql.DB {
	// Capture connection properties.
	connStr := "user=postgres dbname=fitworks sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingDb()
	getCurrentDb()
	return db
}

func pingDb() {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}

func getCurrentDb() {
	var currentDatabase string
	query := "SELECT CURRENT_DATABASE()"
	row := db.QueryRow(query)
	if err := row.Scan(&currentDatabase); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("No current database found")
		}
		log.Fatal(err)
	}
	log.Println("Connected to database: " + currentDatabase)
}
