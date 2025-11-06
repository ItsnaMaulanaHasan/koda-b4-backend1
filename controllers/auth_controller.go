package controllers

import (
	"gin-practice/lib"
	"gin-practice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/matthewhartstonge/argon2"
)

type AuthController struct {
	userController *UserController
}

func NewAuthController(uc *UserController) *AuthController {
	return &AuthController{
		userController: uc,
	}
}

// Register godoc
// @Summary      Create new user
// @Description  Create a new user with a unique username and email
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        user      formData  models.User true "User registration data"
// @Success      200       {object}  models.Response{data=models.User}  "User created successfully"
// @Failure      400       {object}  models.Response  "Invalid request body or hash password failed"
// @Failure      409       {object}  models.Response  "Email or username already exists"
// @Router       /auth/register [post]
func (ac *AuthController) Register(ctx *gin.Context) {
	var body models.User
	err := ctx.ShouldBindWith(&body, binding.Form)
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

	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Hash password failed",
		})
		return
	}

	body.Password = string(hashPassword)
	ac.userController.users = append(ac.userController.users, body)

	body.Password = ""

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    body,
	})
}

// Login godoc
// @Summary      Login user
// @Description  Log in with existing email and username data
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        email     formData  string  true  "Email address"
// @Param        password  formData  string  true  "Input Password" format(password)
// @Success      200       {object}  models.Response{data=models.User}  "User updated successfully"
// @Failure      400       {object}  models.Response  "Invalid Id format or request body"
// @Failure      404       {object}  models.Response  "User not found"
// @Router       /auth/login [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=6"`
	}

	err := ctx.ShouldBindWith(&loginData, binding.Form)
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
			isPasswordValid, err := argon2.VerifyEncoded([]byte(loginData.Password), []byte(ac.userController.users[i].Password))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, models.Response{
					Success: false,
					Message: err.Error(),
				})
			}

			if isPasswordValid {
				foundUser = &ac.userController.users[i]
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

func (ac *AuthController) ForgotPassword(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	var newPassword struct {
		NewPassword string `json:"newPassword" binding:"required,min=6"`
	}

	err = ctx.ShouldBindWith(&newPassword, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	hashPassword, err := lib.HashPassword(newPassword.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Hash password failed",
		})
		return
	}

	var foundUser *models.User
	for i := range ac.userController.users {
		if ac.userController.users[i].Id == id {
			ac.userController.users[i].Password = string(hashPassword)
			foundUser = &ac.userController.users[i]
		}
	}

	if foundUser != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Password success updates",
			Data:    foundUser,
		})
	}
}
