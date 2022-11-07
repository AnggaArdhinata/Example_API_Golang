package interfaces

import (
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
)

type AuthService interface {
	Login(user *models.User) (*helpers.Response, error)
	Register(user *models.User) (*helpers.Response, error)
}
