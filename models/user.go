package models

type User struct {
	Id       int    `json:"id" form:"id" xml:"id"`
	Username string `json:"username" form:"username" xml:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" form:"email" xml:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password" xml:"password" binding:"required,min=6"`
}
