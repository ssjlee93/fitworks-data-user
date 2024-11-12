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
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	var currentDatabase string
	row := db.QueryRow("SELECT CURRENT_DATABASE()")
	if err := row.Scan(&currentDatabase); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("No current database found")
			return nil
		}
		log.Fatal(err)
		return nil
	}
	log.Println("Connected to database: " + currentDatabase)
	return db
}
