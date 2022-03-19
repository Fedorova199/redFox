package handlers

import (
	"io"
	"net/http"

	"github.com/malyg1n/shortener/api/rest/models"
)

func HandlerGet(w http.ResponseWriter, r *http.Request) {

}

func (m *models.Models) HandlerPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}

	url := string(body)

}
