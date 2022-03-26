package storage

import (
	"database/sql"
	"log"

	"github.com/Fedorova199/redfox/internal/config"
)

func OpenDB(cfg config.Config) {
	//dns := "user=postgres password=password dbname=urls sslmode=disable"
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
