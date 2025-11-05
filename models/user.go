package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username,omitempty" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}
