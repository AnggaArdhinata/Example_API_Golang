package auth

import (
	"encoding/json"
	"net/http"
	"restapiexample/src/databases/models"
	"restapiexample/src/interfaces"
)

type AuthController struct {
	Auth_Service interfaces.AuthService
}

func NewController(authservice interfaces.AuthService) *AuthController {
	return &AuthController{authservice}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	result, err := c.Auth_Service.Login(&user)
	if err != nil {
		result.Send(w)
	} else {
		result.Send(w)
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	result, err := c.Auth_Service.Register(&user)
	if err != nil {
		result.Send(w)
	} else {
		result.Send(w)
	}
}