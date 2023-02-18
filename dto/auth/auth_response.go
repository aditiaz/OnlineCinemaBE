package authdto

type RegisterResponse struct {
	Fullname string ` json:"fullname"`
	Message  string ` json:"message"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string ` json:"token"`
}
