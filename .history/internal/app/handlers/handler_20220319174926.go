package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fedorova199/shortURL/internal/app/storage"
	"github.com/go-chi/chi"
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

func (m *Models) HandlerJson(w http.ResponseWriter, r *http.Request) {
	m.Counter++
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

	err = m.AddURL(fmt.Sprintf("%d", m.Counter), urljson.URL)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	result := storage.ResponseJSON{
		Response: "http://localhost:8080/" + fmt.Sprintf("%d", m.Counter),
	}

	response, _ := json.Marshal(result)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
