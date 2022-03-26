package storage

import (
	"database/sql"
	"log"
)

func OpenDB() {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
