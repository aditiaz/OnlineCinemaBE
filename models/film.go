package models

type Film struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	FilmUrl     string  `json:"film_url"`
	Description string  `json:"description"`
	Thumbnail   string  `json:"thumbnail"`
}

func (Film) TableName() string {
	return "films"
}
