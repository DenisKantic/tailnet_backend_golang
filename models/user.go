package models

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Location        string `json:"location" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Bio             string `json:"bio" binding:"required"`
}
