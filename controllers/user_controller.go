package controllers

import (
	"fmt"
	"gin-practice/lib"
	"gin-practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	users []models.User
}

func NewUserController() *UserController {
	return &UserController{
		users: []models.User{},
	}
}

// GetAllUser godoc
// @Summary      Get all users
// @Description  Retrieving all user data
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.Response{data=[]models.User}  "Success get all users"
// @Router       /users [get]
func (uc *UserController) GetAllUser(ctx *gin.Context) {
	var responseData = []models.User{}
	for _, u := range uc.users {
		responseData = append(responseData, models.User{
			Id:       u.Id,
			Username: u.Username,
			Email:    u.Email,
		})
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success get all user",
		Data:    responseData,
	})
}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Retrieving user data based on Id
// @Tags         users
// @Accept 		 x-www-form-urlencoded
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.Response{data=models.User}  "Success get user"
// @Failure      400  {object}  models.Response  "Invalid Id format"
// @Failure      404  {object}  models.Response  "User not found"
// @Router       /users/{id} [get]
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
		foundUser.Password = ""

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

// CreateUser godoc
// @Summary      Create new user
// @Description  Create a new user with a unique username and email
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        user      formData  models.User true "User registration data"
// @Success      200       {object}  models.Response{data=models.User}  "User created successfully"
// @Failure      400       {object}  models.Response  "Invalid request body or hash password failed"
// @Failure      409       {object}  models.Response  "Email or username already exists"
// @Router       /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var body models.User
	err := ctx.ShouldBindWith(&body, binding.Form)
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

	body.Password = ""

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Create user succesfully",
		Data:    body,
	})
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Updating user data (username and email) based on Id
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        id        path      int     true  "User ID"
// @Param        username  formData  string  true  "Username (min 3, max 20 chars)"
// @Param        email     formData  string  true  "Email address"
// @Success      200       {object}  models.Response{data=models.User}  "User updated successfully"
// @Failure      400       {object}  models.Response  "Invalid Id format or request body"
// @Failure      404       {object}  models.Response  "User not found"
// @Router       /users/{id} [patch]
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
		Username string `form:"username" binding:"required,min=3,max=20"`
		Email    string `form:"email" binding:"required,email"`
	}
	err = ctx.ShouldBindWith(&userUpdate, binding.Form)

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
		foundUser.Password = ""
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

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by Id
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        id   path      int  true  "User Id"
// @Success      200  {object}  models.Response{data=models.User}  "User deleted successfully"
// @Failure      400  {object}  models.Response  "Invalid Id format"
// @Failure      404  {object}  models.Response  "User not found"
// @Router       /users/{id} [delete]
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
		foundUser.Password = ""
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
