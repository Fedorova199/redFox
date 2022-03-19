package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"test.txt"`
}

const (
	defaultServerAddress = ":8080"
	defaultBaseURL       = "http://localhost:8080"
)

var defaultConfig = Config{
	ServerAddress: defaultServerAddress,
	BaseURL:       defaultBaseURL,
}

func NewConfig() (Config, error) {
	conf := defaultConfig
	conf.parseFlags()
	conf.parseEnvVars()
	return conf, nil
}

func (conf *Config) parseFlags() {

	flag.StringVar(&conf.ServerAddress, "a", defaultServerAddress, "address")
	flag.StringVar(&conf.BaseURL, "b", defaultBaseURL, "base URL")
	flag.StringVar(&conf.FileStoragePath, "f", "test.txt", "storage file")
	flag.Parse()

}

func (conf *Config) parseEnvVars() {
	sa := os.Getenv("SERVER_ADDRESS")
	if sa != "" {
		conf.ServerAddress = sa
	}

	bu := os.Getenv("BASE_URL")
	if bu != "" {
		conf.BaseURL = bu
	}

	fsp := os.Getenv("FILE_STORAGE_PATH")
	if fsp != "" {

		conf.FileStoragePath = fsp
	}

}
