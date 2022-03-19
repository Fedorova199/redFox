package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/malyg1n/shortener/storage"
)

func (m *Models) HandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	id := chi.URLParam(r, "Id")
	URL, err := m.GetURL(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func (m *Models) HandlerPost(w http.ResponseWriter, r *http.Request) {
	m.Counter++
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}

	//cfg := config.Config{}
	url := string(body)
	fmt.Println(url)
	err = m.AddURL(fmt.Sprintf("%d", m.Counter), url)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	shURL := "http://localhost:8080/" + fmt.Sprintf("%d", m.Counter)
	fmt.Println(shURL)
	w.Write([]byte(shURL))

}

func (m *Models) JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var urljson storage.RequestJSON
	dataBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(dataBytes, &urljson)
	if err != nil || urljson.URL == "" {
		http.Error(w, "invalid parse body", http.StatusBadRequest)
		return
	}

	key := s.service.CreateRedirect(redirect.URL)
	result := ResultString{
		Result: fmt.Sprintf("%s/%s", s.ctx.ServiceURL, key),
	}

	response, _ := json.Marshal(result)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
