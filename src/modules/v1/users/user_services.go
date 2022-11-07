package users

import (
	"fmt"
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
	"restapiexample/src/interfaces"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) GetAll() (*helpers.Response, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	
	response := helpers.New(200, false, "success get all users", users)
	return response, nil
}

func (s *UserService) GetByID(id uint) (*helpers.Response, error) {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}
	user.Password="secret"

	response := helpers.New(200, false, "success get user by id", user)
	return response, nil
}

func (s *UserService) GetByEmail(email string) (*helpers.Response, error) {
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	response := helpers.New(200, false, "success get user by email", user)
	return response, nil
}


func (s *UserService) Add(user *models.User) (*helpers.Response, error) {
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	user.Password = hashPassword
	data, err := s.UserRepository.Add(user)
	fmt.Println(data)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}
	response := helpers.New(201, false, "success add user", data)
	return response, nil
}

func (s *UserService) Update(id uint, user *models.User) (*helpers.Response, error) {
	userCheck, err := s.UserRepository.GetByID(id)
	fmt.Println(userCheck)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	user.Password = hashPassword
	data, err := s.UserRepository.Update(id, user)
	fmt.Println(data)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	

	response := helpers.New(201, false, "success update user", data)
	return response, nil
}
func (s *UserService) Delete(user *models.User) (*helpers.Response, error) {
	user, err := s.UserRepository.GetByID(user.ID)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	data, err := s.UserRepository.Delete(user)
	if err != nil {
		response := helpers.New(500, true, err.Error(), nil)
		return response, err
	}

	response := helpers.New(200, false, "success delete user", data)
	return response, nil
}
