package products

import (
	"restapiexample/src/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/products").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/q", controller.GetByName).Methods("GET")
	route.HandleFunc("/{id}", middlewares.Chain(controller.GetByID, middlewares.AuthMiddleware)).Methods("GET")
	route.HandleFunc("/", middlewares.Chain(controller.Add, middlewares.AuthMiddleware, middlewares.IsAdmin)).Methods("POST")
	route.HandleFunc("/{id}", middlewares.Chain(controller.Update, middlewares.AuthMiddleware, middlewares.IsAdmin)).Methods("PUT")
	route.HandleFunc("/{id}", middlewares.Chain(controller.Delete, middlewares.AuthMiddleware, middlewares.IsAdmin)).Methods("DELETE")
}
