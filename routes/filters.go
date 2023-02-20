package routes

import (
	"finaltaskbe/handlers"
	"finaltaskbe/pkg/mysql"
	"finaltaskbe/repositories"

	"github.com/gorilla/mux"
)

func FilterRoutes(r *mux.Router) {
	filterRepository := repositories.RepositoryFilter(mysql.DB)
	h := handlers.HandlerFilter(filterRepository)

	r.HandleFunc("/filtercategory", h.MultiFilter).Methods("GET")

}
