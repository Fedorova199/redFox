package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/Fedorova199/redfox/internal/config"
	"github.com/Fedorova199/redfox/internal/handlers"
	"github.com/Fedorova199/redfox/internal/middlewares"
	"github.com/Fedorova199/redfox/internal/storage"
)

func main() {
	cfg := config.ParseVariables()
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	storage, err := storage.CreateDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	ms := []handlers.Middleware{
		middlewares.GzipHandle{},
		middlewares.UngzipHandle{},
		middlewares.NewAuthenticator([]byte("secret key")),
	}

	handler := handlers.NewHandler(storage, cfg.BaseURL, ms)
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: handler,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-c
		server.Close()
	}()

	log.Fatal(server.ListenAndServe())
}
