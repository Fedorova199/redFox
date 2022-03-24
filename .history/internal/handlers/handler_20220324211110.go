package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (md *Models) POSTHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	longURL := string(body)
	if longURL == " " {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	md.counter++
	md.model[fmt.Sprintf("%d", md.counter)] = longURL
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
}

func (md *Models) GETHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id != fmt.Sprintf("%d", md.counter) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, md.model[id], http.StatusTemporaryRedirect)
}

func (md *Models) JSONHandler(w http.ResponseWriter, r *http.Request) {
	md.counter++
	body, _ := ioutil.ReadAll(r.Body)
	someURL := URLRequest{}
	if err := json.Unmarshal([]byte(body), &someURL); err != nil {
		log.Fatalln(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	md.model[fmt.Sprintf("%d", md.counter)] = someURL.SomeURL

	shortenerURL := URLResponse{
		ShortenerURL: "http://localhost:8080/" + fmt.Sprintf("%d", md.counter),
	}
	js, err := json.MarshalIndent(&shortenerURL, " ", "")
	if err != nil {
		log.Fatalln(err)
		return
	}
	w.Write(js)

}
