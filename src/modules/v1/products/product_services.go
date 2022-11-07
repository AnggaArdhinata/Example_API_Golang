package products

import (
	"fmt"
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
	"restapiexample/src/interfaces"
)

type ProductService struct {
	ProductRepository interfaces.ProductRepository
}

func NewService(productRepository interfaces.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (s *ProductService) GetAll() (*helpers.Response, error) {
	products, err := s.ProductRepository.GetAll()
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	response := helpers.New(200, false, "success get all products", products)
	return response, nil
}

func (s *ProductService) GetByID(id uint) (*helpers.Response, error) {
	product, err := s.ProductRepository.GetByID(id)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	response := helpers.New(200, false, "success get product by id", product)
	return response, nil
}

func (s *ProductService) GetByName(name string) (*helpers.Response, error) {
	product, err := s.ProductRepository.GetByName(name)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}
	response := helpers.New(200, false, "success get product by name", product)
	return response, nil
}

func (s *ProductService) Add(product *models.Product) (*helpers.Response, error) {
	product, err := s.ProductRepository.Add(product)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}
	response := helpers.New(200, false, "success add products", product)
	return response, nil

}

func (s *ProductService) Update(id uint, product *models.Product) (*helpers.Response, error) {
	productCheck, err := s.ProductRepository.GetByID(id)
	fmt.Println(productCheck)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	product, err = s.ProductRepository.Update(id, product)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}
	response := helpers.New(200, false, "success edit product", product)
	return response, nil
}

func (s *ProductService) Delete(product *models.Product) (*helpers.Response, error) {
	productCheck, err := s.ProductRepository.GetByID(product.ID)
	fmt.Println(productCheck)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	data, err := s.ProductRepository.Delete(product)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	response := helpers.New(200, false, "success delete product", data)
	return response, nil
}
