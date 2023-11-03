package db

import (
	"database/sql"
	"log"
)

func ConectaComBancoDeDados() *sql.DB {
	connection := "user=alessandro-maciel dbname=alessandro-maciel host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
