package users

import (
	"restapiexample/src/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(router *mux.Router, db *gorm.DB) {
	route := router.PathPrefix("/users").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/q", middlewares.Chain(controller.GetByEmail, middlewares.AuthMiddleware, middlewares.IsAdmin)).Methods("GET")
	route.HandleFunc("/{id}", controller.GetByID).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")
	route.HandleFunc("/{id}", middlewares.Chain(controller.Update, middlewares.AuthMiddleware)).Methods("PUT")
	route.HandleFunc("/{id}", middlewares.Chain(controller.Delete, middlewares.AuthMiddleware, middlewares.IsAdmin)).Methods("DELETE")
}
