package filmsdto

type FilmResponse struct {
	Title       string  `json:"title"`
	Category    string  `json:"category"`
	Price       float64 `json"price"`
	FilmUrl     string  `json "film_url"`
	Description string  `json:"description"`
}
