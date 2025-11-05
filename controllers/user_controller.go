package controllers

import (
	"fmt"
	"gin-practice/lib"
	"gin-practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	users []models.User
}

func NewUserController() *UserController {
	return &UserController{
		users: []models.User{},
	}
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success get all user",
		Data:    uc.users,
	})
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id format",
		})
		return
	}

	var foundUser *models.User
	for i := range uc.users {
		if uc.users[i].Id == id {
			foundUser = &uc.users[i]
		}
	}

	if foundUser != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("Success get user with id %d", id),
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: fmt.Sprintf("User with id %d not found", id),
		})
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var body models.User
	err := ctx.ShouldBindBodyWithJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	for _, user := range uc.users {
		if user.Email == body.Email {
			ctx.JSON(http.StatusConflict, models.Response{
				Success: false,
				Message: "Email already registered",
			})
			return
		}
	}

	for _, user := range uc.users {
		if user.Username == body.Username {
			ctx.JSON(http.StatusConflict, models.Response{
				Success: false,
				Message: "Username already taken",
			})
			return
		}
	}

	body.Id = len(uc.users) + 1

	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Hash password failed",
		})
		return
	}

	body.Password = string(hashPassword)
	uc.users = append(uc.users, body)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Create user succesfully",
		Data:    body,
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id format",
		})
		return
	}

	var userUpdate struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
	}
	err = ctx.ShouldBindBodyWithJSON(&userUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	var foundUser *models.User
	for i := range uc.users {
		if uc.users[i].Id == id {
			uc.users[i].Username = userUpdate.Username
			uc.users[i].Email = userUpdate.Email
			foundUser = &uc.users[i]
		}
	}

	if foundUser != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("User with id %d successfully updated", id),
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: fmt.Sprintf("User with id %d not found", id),
		})
	}
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id format",
		})
		return
	}

	var foundUser *models.User
	for i := range uc.users {
		if uc.users[i].Id == id {
			foundUser = &uc.users[i]
			uc.users = append(uc.users[:i], uc.users[i+1:]...)
			break
		}
	}

	if foundUser != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("User data with id %d successfully deleted", id),
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: fmt.Sprintf("User with id %d not found", id),
		})
	}
}
