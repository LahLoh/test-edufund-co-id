package model

type RegisterRequest struct {
	Fullname             string `json:"fullname"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmation_password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
