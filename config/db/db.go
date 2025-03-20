package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=streaming sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, nil
}
