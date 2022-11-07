package routers

import (
	"net/http"
	database "restapiexample/src/databases"
	"restapiexample/src/modules/v1/auth"
	"restapiexample/src/modules/v1/users"
	"restapiexample/src/modules/v1/products"
	


	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	subRouter := mainRoute.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/", exampleHandler).Methods("GET")

	auth.New(subRouter, db)
	users.New(subRouter, db)
	products.New(subRouter, db)
	
	return subRouter, nil
}


func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
