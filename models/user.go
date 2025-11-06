package models

type User struct {
	Id           int    `form:"id" swaggerignore:"true"`
	Username     string `form:"username" binding:"required,min=3,max=20" example:"koda"`
	Email        string `form:"email" binding:"required,email" example:"koda@mail.com"`
	Password     string `form:"password" json:"password,omitempty" binding:"required,min=6" example:"koda123" format:"password"`
	PhotoProfile string `form:"profilePhoto,omitempty" swaggerignore:"true" json:"profilePhoto,omitempty"`
}
