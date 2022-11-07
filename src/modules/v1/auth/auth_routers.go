package auth

import (

	"restapiexample/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(router *mux.Router, db *gorm.DB) {
	route := router.PathPrefix("/auth").Subrouter()

	repository := users.NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/login", controller.Login).Methods("POST")
	route.HandleFunc("/register", controller.Register).Methods("POST")
}