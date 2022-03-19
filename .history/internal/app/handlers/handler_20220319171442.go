package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func (m Models) HandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	id := chi.URLParam(r, "Id")
	URL, err := m.GetURL(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func (m Models) HandlerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}

	//cfg := config.Config{}
	url := string(body)
	fmt.Println(url)
	err = m.AddURL(url)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	shURL := "http://localhost:8080/" + fmt.Sprintf("%d", m.Counter+1)
	fmt.Println(shURL)
	w.Write([]byte(shURL))

}
