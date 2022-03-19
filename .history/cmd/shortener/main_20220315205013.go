package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Fedorova199/shortURL/internal/app/config"
	"github.com/Fedorova199/shortURL/internal/app/server"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGTERM, syscall.SIGINT,
	)
	defer stop()
	var cfg config.Config
	router := server.NewRouter()
	server, err := server.NewHTTPServer(cfg.ServerAddress, router)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start(ctx)
	}()

	wg.Wait()
	if err := store.Close(); err != nil {
		log.Fatalln(err)
	}

}
