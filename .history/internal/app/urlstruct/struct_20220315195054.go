package structs

type Models struct {
	ShortURL map[string]string
}

func NewModels() *Models {
	model := Models{
		ShortURL: make(map[string]string),
	}
	return &model
}
