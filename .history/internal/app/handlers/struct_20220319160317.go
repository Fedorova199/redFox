package handlers

import "fmt"

type Models struct {
	ShortURL map[string]string
}

func NewModels() *Models {
	model := Models{
		ShortURL: make(map[string]string),
	}
	return &model
}

func (m *Models) GetURL(id string) (string, error) {
	if id == "" {
		return "", fmt.Errorf("empty id")
	}
	return m.ShortURL[id], nil
}

func (m *Models) AddURL(id string, url string) {
	m.ShortURL[id] = url
}
