package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	FilmRoutes(r)
	// TransactionRoutes(r)
	// FilterRoutes(r)
}
