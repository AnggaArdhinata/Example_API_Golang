package interfaces

import (
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
)

type UserRepository interface {
	GetAll() (*models.Users, error)
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Add(user *models.User) (*models.User, error)
	Update(id uint, user *models.User) (*models.User, error)
	Delete(user *models.User) (uint, error)
}

type UserService interface {
	GetAll() (*helpers.Response, error)
	GetByID(id uint) (*helpers.Response, error)
	GetByEmail(email string) (*helpers.Response, error)
	Add(user *models.User) (*helpers.Response, error)
	Update(id uint, user *models.User) (*helpers.Response, error)
	Delete(user *models.User) (*helpers.Response, error)
}
