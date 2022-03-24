package handlers

type Models struct {
	model   map[string]string
	counter int
}

type URLRequest struct {
	SomeURL string `json:"url"`
}

type URLResponse struct {
	ShortenerURL string `json:"results"`
}

func NewModels() *Models {
	return &Models{
		model:   make(map[string]string),
		counter: 0,
	}
}
