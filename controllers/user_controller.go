package controllers

import (
	"fmt"
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
		users: []models.User{
			{Id: 1, Username: "Itsna", Email: "itsna@mail.com", Password: "123456789"},
			{Id: 2, Username: "Federus", Email: "federus@mail.com", Password: "123456789"},
			{Id: 3, Username: "Ari", Email: "ari@mail.com", Password: "123456789"},
			{Id: 4, Username: "Yoga", Email: "yoga@mail.com", Password: "123456789"},
			{Id: 5, Username: "Fiki", Email: "fiki@mail.com", Password: "123456789"},
			{Id: 6, Username: "Sidik", Email: "sidik@mail.com", Password: "123456789"},
			{Id: 7, Username: "Anggi", Email: "anggi@mail.com", Password: "123456789"},
		},
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
			Message: "Invliad Id format",
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

	err := ctx.ShouldBind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	uc.users = append(uc.users, body)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success add data user",
		Data:    body,
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invliad Id format",
		})
		return
	}

	var body models.User
	err = ctx.ShouldBind(&body)

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
			uc.users[i].Username = body.Username
			uc.users[i].Email = body.Email
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
			Message: "Invliad Id format",
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
