package handlers

import (
	"io"
	"net/http"
)

func HandlerGet(w http.ResponseWriter, r *http.Request) {

}

func (m Models) HandlerPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Server failed to read the request's body", http.StatusInternalServerError)
		return
	}

	url := string(body)

}
