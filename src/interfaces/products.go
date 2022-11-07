package interfaces

import (
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
)

type ProductRepository interface {
	GetAll() (*models.Products, error)
	GetByID(id uint) (*models.Product, error)
	GetByName(name string) (*models.Product, error)
	Add(*models.Product) (*models.Product, error)
	Update(id uint, data *models.Product) (*models.Product, error)
	Delete(data *models.Product) (uint, error)
}

type ProductService interface {
	GetAll() (*helpers.Response, error)
	GetByID(id uint) (*helpers.Response, error)
	GetByName(name string) (*helpers.Response, error)
	Add(*models.Product) (*helpers.Response, error)
	Update(id uint, data *models.Product) (*helpers.Response, error)
	Delete(data *models.Product) (*helpers.Response, error)
}