package auth

import (
	"fmt"
	"restapiexample/src/databases/models"
	"restapiexample/src/helpers"
	"restapiexample/src/interfaces"
)

type Auth_Service struct {
	UserRepository interfaces.UserRepository
}

func NewService(userRepository interfaces.UserRepository) *Auth_Service {
	return &Auth_Service{userRepository}
}

func (s *Auth_Service) Login(user *models.User) (*helpers.Response, error) {
	checkEmail, err := s.UserRepository.GetByEmail(user.Email)
	if err != nil {
		response := helpers.New(500, true, "failed get user by email", nil)
		return response, err
	}

	if !helpers.ComparePassword(user.Password, checkEmail.Password) {
		response := helpers.New(400, true, "password invalid", nil)
		return response, nil
	}

	token, err := helpers.GenerateToken(checkEmail.Name, checkEmail.Role)
	if err != nil {
		response := helpers.New(500, true, "failed generate token", nil)
		return response, err
	}

	data := LoginResponse{
		JWT_Token: token,
		
	}
	response := helpers.New(200, false, "login success", data)
	return response, nil

}

func (s *Auth_Service) Register(user *models.User) (*helpers.Response, error) {
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

	dataResponse := RegisterResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	response := helpers.New(201, false, "register success", dataResponse)
	return response, nil
}
