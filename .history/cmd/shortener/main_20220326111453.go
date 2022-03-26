package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fedorova199/shorturl/internal/config"
	"github.com/Fedorova199/shorturl/internal/handlers"
	"github.com/go-chi/chi"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	model := handlers.NewModels()
	router := chi.NewRouter()
	fmt.Println(cfg)
	router.Post("/", model.POSTHandler)
	router.Post("/api/shorten", model.JSONHandler)
	router.Get("/{id}", model.GETHandler)

	http.ListenAndServe(cfg.ServerAddress, router)

}
