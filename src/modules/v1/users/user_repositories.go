package users

import (
	"errors"
	"fmt"
	"restapiexample/src/databases/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(gorm *gorm.DB) *UserRepository {
	return &UserRepository{gorm}
}

func (r *UserRepository) GetAll() (*models.Users, error) {
	var users models.Users

	result := r.db.Order("created_at desc").Find(&users)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}


func (r *UserRepository) Add(user *models.User) (*models.User, error) {
	emailCheck := r.db.First(&user, "email = ?", user.Email)
	if emailCheck.Error == nil {
		return nil, errors.New("email already exist")
	}

	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) Update(id uint, user *models.User) (*models.User, error) {
	result := r.db.Model(user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) Delete(user *models.User) (uint, error) {
	result := r.db.Where("id = ?", user.ID).Delete(user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
