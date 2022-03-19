package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Fedorova199/shortURL/internal/app/handlers"
	"github.com/go-chi/chi"
	"honnef.co/go/tools/config"
)

type HTTPServer struct {
	srv             *http.Server
	idleConnsClosed chan struct{}
	shutdownTimeout time.Duration
}

func NewRouter(cxt context.Context, cfg config.Config) http.Handler {
	router := chi.NewRouter()

	urlHandler := handlers.NewModels()

	router.Get("/{id}", urlHandler.HandlerGet)
	router.Post("/", urlHandler.HandlerPost)
	//router.Post("/api/shorten", urlHandler.JSONHandler)
	return router
}

func NewHTTPServer(address string, router http.Handler) (*HTTPServer, error) {

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	return &HTTPServer{
		srv:             server,
		idleConnsClosed: make(chan struct{}),
		shutdownTimeout: 3 * time.Second,
	}, nil
}

func (s *HTTPServer) Close() {
	if s.srv == nil {
		return
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	if err := s.srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalln(err)
		return
	}

	s.srv = nil
	close(s.idleConnsClosed)
}

func (s *HTTPServer) Start(ctx context.Context) {
	if s.srv == nil {
		return
	}
	go func() {
		<-ctx.Done()
		s.Close()
	}()

	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
		return
	}

	<-s.idleConnsClosed
}
