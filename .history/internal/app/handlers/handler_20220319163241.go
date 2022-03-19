package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Fedorova199/shortURL/internal/app/config"
)

func (m Models) HandlerGet(w http.ResponseWriter, r *http.Request) {

}

func (m Models) HandlerPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}
	cfg := config.Config{}
	url := string(body)
	m.AddURL(url)
	shURL := cfg.BaseURL + fmt.Sprintf("%d", m.Counter)
	w.Write([]byte(shURL))

}
