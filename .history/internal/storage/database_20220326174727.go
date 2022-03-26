package storage

import (
	"database/sql"
	"log"

	"github.com/Fedorova199/redfox/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenDB(cfg config.Config) {
	//dns := "user=postgres password=password dbname=urls sslmode=disable"
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
