package routes

import (
	"finaltaskbe/handlers"
	"finaltaskbe/pkg/middleware"
	"finaltaskbe/pkg/mysql"
	"finaltaskbe/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/films", h.FindFilms).Methods("GET")
	r.HandleFunc("/publicfilm/{id}", h.GetPublicFilm).Methods("GET")
	r.HandleFunc("/film/{id}", middleware.Auth(h.GetFilm)).Methods("GET")
	r.HandleFunc("/addfilm", middleware.Auth(middleware.UploadFile(h.Addfilm))).Methods("POST")

}
