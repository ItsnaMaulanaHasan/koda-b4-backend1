package controllers

import (
	"gin-practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userController *UserController
}

func NewAuthController(uc *UserController) *AuthController {
	return &AuthController{
		userController: uc,
	}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var body models.User
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	for _, user := range ac.userController.users {
		if user.Email == body.Email {
			ctx.JSON(http.StatusConflict, models.Response{
				Success: false,
				Message: "Email already registered",
			})
			return
		}
	}

	for _, user := range ac.userController.users {
		if user.Username == body.Username {
			ctx.JSON(http.StatusConflict, models.Response{
				Success: false,
				Message: "Username already taken",
			})
			return
		}
	}

	body.Id = len(ac.userController.users) + 1

	ac.userController.users = append(ac.userController.users, body)

	responseData := models.User{
		Id:       body.Id,
		Username: body.Username,
		Email:    body.Email,
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    responseData,
	})
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" form:"email" xml:"email" binding:"required,email"`
		Password string `json:"password" form:"password" xml:"password" binding:"required,min=6"`
	}

	err := ctx.ShouldBind(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var foundUser *models.User
	for i := range ac.userController.users {
		if ac.userController.users[i].Email == loginData.Email {
			if ac.userController.users[i].Password == loginData.Password {
				foundUser = &ac.userController.users[i]
				break
			} else {
				ctx.JSON(http.StatusUnauthorized, models.Response{
					Success: false,
					Message: "Invalid email or password",
				})
				return
			}
		}
	}

	if foundUser == nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	responseData := models.User{
		Id:       foundUser.Id,
		Username: foundUser.Username,
		Email:    foundUser.Email,
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User login Successfully",
		Data:    responseData,
	})
}
