package controllers

import (
	"fmt"
	"gin-practice/lib"
	"gin-practice/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
// @Description  Retrieving all user data with pagination support
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization    header    string  true  "Bearer token"  default(Bearer <token>)
// @Param        page   query     int  false  "Page number"  default(1)  minimum(1)
// @Param        limit  query     int  false  "Number of items per page"  default(10)  minimum(1)  maximum(100)
// @Success      200    {object}  object{success=bool,message=string,data=[]models.User,meta=object{currentPage=int,perPage=int,totalData=int,totalPages=int}}  "Success get all users"
// @Failure      400    {object}  lib.Response  "Invalid pagination parameters"
// @Router       /users [get]
func (uc *UserController) GetAllUser(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page < 1 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Page must be greater than 0",
		})
		return
	}

	if limit < 1 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Limit must be greater than 0",
		})
		return
	}

	if limit > 100 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Limit cannot exceed 100",
		})
		return
	}

	totalData := len(uc.users)
	totalPage := (totalData + limit - 1) / limit
	startIndex := (page - 1) * limit
	endIndex := min(startIndex+limit, totalData)

	var responseData []models.User

	if startIndex >= totalData && totalData > 0 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Page is out of range",
		})
		return
	}

	if startIndex < totalData {
		for _, u := range uc.users[startIndex:endIndex] {
			responseData = append(responseData, models.User{
				Id:           u.Id,
				Username:     u.Username,
				Email:        u.Email,
				PhotoProfile: u.PhotoProfile,
			})
		}
	} else {
		responseData = []models.User{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success get all user",
		"data":    responseData,
		"meta": gin.H{
			"currentPage": page,
			"perPage":     limit,
			"totalData":   totalData,
			"totalPages":  totalPage,
		},
	})
}

// GetUserById godoc
// @Summary      Get user by Id
// @Description  Retrieving user data based on Id
// @Tags         users
// @Accept 		 x-www-form-urlencoded
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Bearer token"  default(Bearer <token>)
// @Param        id   path      int  true  "User Id"
// @Success      200  {object}  lib.Response{data=models.User}  "Success get user"
// @Failure      400  {object}  lib.Response  "Invalid Id format"
// @Failure      404  {object}  lib.Response  "User not found"
// @Router       /users/{id} [get]
func (uc *UserController) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
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

		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Success get user",
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
	}
}

// CreateUser godoc
// @Summary      Create new user
// @Description  Create a new user with a unique username and email
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Bearer token"  default(Bearer <token>)
// @Param        user      formData  models.User true "User registration data"
// @Success      200       {object}  lib.Response{data=models.User}  "User created successfully"
// @Failure      400       {object}  lib.Response  "Invalid request body or hash password failed"
// @Failure      409       {object}  lib.Response  "Email or username already exists"
// @Router       /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var body models.User
	err := ctx.ShouldBindWith(&body, binding.Form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: true,
			Message: err.Error(),
		})
		return
	}

	for _, user := range uc.users {
		if user.Email == body.Email {
			ctx.JSON(http.StatusConflict, lib.Response{
				Success: false,
				Message: "Email already registered",
			})
			return
		}
	}

	for _, user := range uc.users {
		if user.Username == body.Username {
			ctx.JSON(http.StatusConflict, lib.Response{
				Success: false,
				Message: "Username already taken",
			})
			return
		}
	}

	body.Id = len(uc.users) + 1

	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Hash password failed",
		})
		return
	}

	body.Password = string(hashPassword)
	uc.users = append(uc.users, body)

	body.Password = ""

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "User created successfully",
		Data:    body,
	})
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Updating user data (username and email) based on Id
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Bearer token"  default(Bearer <token>)
// @Param        id        path      int     true  "User Id"
// @Param        username  formData  string  true  "Username (min 3, max 20 chars)"
// @Param        email     formData  string  true  "Email address"
// @Success      200       {object}  lib.Response{data=models.User}  "User updated successfully"
// @Failure      400       {object}  lib.Response  "Invalid Id format or request body"
// @Failure      404       {object}  lib.Response  "User not found"
// @Router       /users/{id} [patch]
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
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
		ctx.JSON(http.StatusBadRequest, lib.Response{
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
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "User updated successfully",
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
	}
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by Id
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Bearer token"  default(Bearer <token>)
// @Param        id   path      int  true  "User Id"
// @Success      200  {object}  lib.Response{data=models.User}  "User deleted successfully"
// @Failure      400  {object}  lib.Response  "Invalid Id format"
// @Failure      404  {object}  lib.Response  "User not found"
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
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
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "User deleted successfully",
			Data:    foundUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
	}
}

// UploadProfile godoc
// @Summary Upload user profile picture
// @Description Upload or replace the profile picture for a user by Id.
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Security     BearerAuth
// @Param        Authorization  header    string  true  "Bearer token"  default(Bearer <token>)
// @Param id path int true "User Id"
// @Param file formData file true "Profile picture (JPEG or PNG, max 1MB)"
// @Success 200 {object} lib.Response "Successfully uploaded photo profile"
// @Failure 400 {object} lib.Response "Bad request (invalid Id, wrong file type, or file too large)"
// @Failure 404 {object} lib.Response "User not found"
// @Failure 500 {object} lib.Response "Failed to save file"
// @Router /users/{id}/upload-profile [patch]
func (uc *UserController) UploadProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var foundUser *models.User
	for i := range uc.users {
		if uc.users[i].Id == id {
			foundUser = &uc.users[i]
			break
		}
	}

	if foundUser == nil {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if file.Size > 1<<20 {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "File size must be less than 1MB",
		})
		return
	}

	contentType := file.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Only image files (JPEG, PNG) are allowed",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("user_%d_%d%s", id, time.Now().Unix(), ext)
	filepath := "./uploads/profiles/" + filename

	if foundUser.PhotoProfile != "" {
		err = os.Remove("./uploads/profiles/" + foundUser.PhotoProfile)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, lib.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	foundUser.PhotoProfile = filename

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Successfully uploaded photo profile",
		Data: map[string]string{
			"filename": filename,
			"url":      "/uploads/profiles/" + filename,
		},
	})
}
