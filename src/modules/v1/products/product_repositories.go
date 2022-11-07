package products

import (
	"errors"
	"fmt"
	"restapiexample/src/databases/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewRepository(gorm *gorm.DB) *ProductRepository {
	return &ProductRepository{gorm}
}

func (r *ProductRepository) GetAll() (*models.Products, error) {
	var Products models.Products

	result := r.db.Order("created_at desc").Find(&Products)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Products, nil
}


func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var Product models.Product

	result := r.db.First(&Product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Product, nil
}

func (r *ProductRepository) GetByName(name string) (*models.Product, error) {
	var Product models.Product

	result := r.db.First(&Product, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Product, nil
}

func (r *ProductRepository) Add(Product *models.Product) (*models.Product, error) {
	emailCheck := r.db.First(&Product, "name = ?", Product.Name)
	if emailCheck.Error == nil {
		return nil, errors.New("name already exist")
	}

	result := r.db.Create(Product)
	if result.Error != nil {
		return nil, result.Error
	}

	return Product, nil
}


func (r *ProductRepository) Update(id uint, Product *models.Product) (*models.Product, error) {
	result := r.db.Model(Product).Where("id = ?", id).Updates(Product)
	if result.Error != nil {
		return nil, result.Error
	}

	return Product, nil
}

func (r *ProductRepository) Delete(Product *models.Product) (uint, error) {
	result := r.db.Where("id = ?", Product.ID).Delete(Product)
	if result.Error != nil {
		return 0, result.Error
	}

	return Product.ID, nil
}