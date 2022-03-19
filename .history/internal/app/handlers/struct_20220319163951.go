package handlers

import "fmt"

type Models struct {
	Counter  int
	ShortURL map[string]string
}

func NewModels() *Models {
	model := Models{
		Counter:  0,
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

func (m *Models) AddURL(url string) {
	id := fmt.Sprintf("%d", m.Counter+1)
	m.ShortURL[id] = url
}
