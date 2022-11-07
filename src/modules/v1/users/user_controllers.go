package users

import (
	"encoding/json"
	"net/http"
	"restapiexample/src/databases/models"
	"restapiexample/src/interfaces"
	"strconv"

	"github.com/gorilla/mux"
)



type UserController struct {
	UserService interfaces.UserService
}

func NewController(userService interfaces.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.UserService.GetAll()
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idToInt, _ := strconv.Atoi(id)
	idToUInt := uint(idToInt)

	response, err := c.UserService.GetByID(idToUInt)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *UserController) GetByEmail(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	email := query.Get("email")
	response, err := c.UserService.GetByEmail(email)

	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}


func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	response, err := c.UserService.Add(&user)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	idToInt, _ := strconv.Atoi(id)
	userID := uint(idToInt)
	user.ID = uint(idToInt)

	json.NewDecoder(r.Body).Decode(&user)
	response, err := c.UserService.Update(userID, &user)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	idToInt, _ := strconv.Atoi(id)
	user.ID = uint(idToInt)

	json.NewDecoder(r.Body).Decode(&user)
	response, err := c.UserService.Delete(&user)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}
