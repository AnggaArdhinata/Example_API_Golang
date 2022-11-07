package auth

type LoginResponse struct {
	JWT_Token  string `json:"JWT_Token"`
	
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}