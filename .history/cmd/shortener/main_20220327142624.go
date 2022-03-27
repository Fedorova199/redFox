package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Fedorova199/redfox/internal/app/handlers"
	"github.com/Fedorova199/redfox/internal/app/interfaces"
	"github.com/Fedorova199/redfox/internal/app/middlewares"
	"github.com/Fedorova199/redfox/internal/app/storage"
	"github.com/caarlos0/env"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

func main() {
	cfg := parseVariables()
	//cfg, err := config.NewConfig()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	storage, err := storage.CreateDatabase(db)
	if err != nil {
		log.Fatalln(err)
	}

	mws := []interfaces.Middleware{
		middlewares.GzipEncoder{},
		middlewares.GzipDecoder{},
		middlewares.NewAuth([]byte("secret key")),
	}

	handler := handlers.NewHandler(storage, cfg.BaseURL, mws)
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

func parseVariables() Config {
	var cfg = Config{
		ServerAddress:   ":8080",
		BaseURL:         "http://localhost:8080",
		FileStoragePath: "test.txt",
	}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File storage path")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "Database DSN")
	flag.Parse()

	return cfg
}
