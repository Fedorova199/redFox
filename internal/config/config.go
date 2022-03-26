package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"test.txt"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

func ParseVariables() Config {
	var cfg = Config{
		ServerAddress:   "localhost:8080",
		BaseURL:         "http://localhost:8080",
		FileStoragePath: "test.txt",
	}

	// err := env.Parse(&cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	sa := os.Getenv("SERVER_ADDRESS")
	if sa != "" {
		cfg.ServerAddress = sa
	}

	bu := os.Getenv("BASE_URL")
	if bu != "" {
		cfg.BaseURL = bu
	}

	fsp := os.Getenv("FILE_STORAGE_PATH")
	if fsp != "" {

		cfg.FileStoragePath = fsp
	}

	dd, ok := os.LookupEnv("DATABASE_DSN")
	if ok {

		cfg.DatabaseDSN = dd
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", `database dsn (default "")`)
	flag.Parse()

	return cfg
}
