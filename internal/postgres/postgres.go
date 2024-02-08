package postgres

import (
	"database/sql"
	"log"
)

func ConnectToPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=1337 host=localhost sslmode=disable port=5433")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db, nil
}
