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

var DataUsers = &UserController{
	users: []models.User{
		{Id: 1, Username: "Itsna", Email: "itsna@mail.com", Password: "12345"},
		{Id: 2, Username: "Federus", Email: "federus@mail.com", Password: "12345"},
		{Id: 3, Username: "Ari", Email: "ari@mail.com", Password: "12345"},
		{Id: 4, Username: "Yoga", Email: "yoga@mail.com", Password: "12345"},
		{Id: 5, Username: "Fiki", Email: "fiki@mail.com", Password: "12345"},
		{Id: 6, Username: "Sidik", Email: "sidik@mail.com", Password: "12345"},
		{Id: 7, Username: "Anggi", Email: "anggi@mail.com", Password: "12345"},
	},
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

	var result models.User
	found := false
	for i := range uc.users {
		if uc.users[i].Id == id {
			result = uc.users[i]
			found = true
		}
	}

	if found {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("Success get user with id %d", id),
			Data:    result,
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

	var result models.User
	found := false
	for i := range uc.users {
		if uc.users[i].Id == id {
			uc.users[i].Username = body.Username
			uc.users[i].Email = body.Email
			result = uc.users[i]
			found = true
		}
	}

	if found {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("User with id %d successfully updated", id),
			Data:    result,
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

	var result models.User
	found := false
	for i := range uc.users {
		if uc.users[i].Id == id {
			result = uc.users[i]
			uc.users = append(uc.users[:i], uc.users[i+1:]...)
			found = true
			break
		}
	}

	if found {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: fmt.Sprintf("User data with id %d successfully deleted", id),
			Data:    result,
		})
	} else {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: fmt.Sprintf("User with id %d not found", id),
		})
	}
}
