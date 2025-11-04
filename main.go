package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var users = []User{
	{
		Id:       1,
		Username: "Itsna",
		Email:    "itsna@mail.com",
		Password: "12345",
	},
	{
		Id:       2,
		Username: "Federus",
		Email:    "federus@mail.com",
		Password: "12345",
	},
	{
		Id:       3,
		Username: "Ari",
		Email:    "ari@mail.com",
		Password: "12345",
	},
	{
		Id:       4,
		Username: "Yoga",
		Email:    "yoga@mail.com",
		Password: "12345",
	},
	{
		Id:       5,
		Username: "Fiki",
		Email:    "fiki@mail.com",
		Password: "12345",
	},
	{
		Id:       6,
		Username: "Sidik",
		Email:    "sidik@mail.com",
		Password: "12345",
	},
	{
		Id:       7,
		Username: "Anggi",
		Email:    "anggi@mail.com",
		Password: "12345",
	},
}

func main() {
	r := gin.Default()

	// Mendapatkan semua data user
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "Success get all user",
			Data:    users,
		})
	})

	// Mendapatkan data user berdasarkan id
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		var result User

		found := false
		for i := range users {
			if users[i].Id == idInt {
				result = users[i]
				found = true
			}
		}

		if found {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("Success get user with id %d", idInt),
				Data:    result,
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: fmt.Sprintf("User with id %d not found", idInt),
			})
		}
	})

	// Menambah user baru
	r.POST("/users", func(ctx *gin.Context) {
		var body User

		err := ctx.BindJSON(&body)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("Failed to add user: %v", err.Error()),
			})
			return
		}

		ctx.JSON(200, Response{
			Success: true,
			Message: "Success add data user",
			Data:    body,
		})

		users = append(users, body)
	})

	// Mengedit data user berdasarkan id
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		var body User

		err := ctx.BindJSON(&body)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("Failed to add user: %v", err.Error()),
			})
			return
		}

		var result User

		found := false
		for i := range users {
			if users[i].Id == idInt {
				users[i].Username = body.Username
				users[i].Email = body.Email
				result = users[i]
				found = true
			}
		}

		if found {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("User with id %d successfully updated", idInt),
				Data:    result,
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User data not found",
			})
		}
	})

	// Menghapus data user by id
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		var result User

		found := false
		for i := range users {
			if users[i].Id == idInt {
				result = users[i]
				users = append(users[:i], users[i+1:]...)
				found = true
				break
			}
		}

		if found {
			ctx.JSON(200, Response{
				Success: true,
				Message: fmt.Sprintf("User data with id %d successfully deleted", idInt),
				Data:    result,
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: "User data not found",
			})
		}
	})

	// r.GET("/", func(ctx *gin.Context) {
	// 	// ctx.Data(200, "application/json", []byte("{\"success\": true, \"message\": \"OK\"}"))

	// 	var data Query

	// 	err := ctx.BindQuery(&data)
	// 	if err != nil {
	// 		ctx.JSON(400, Response{
	// 			Success: false,
	// 			Message: "Error",
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(200, Response{
	// 		Success: true,
	// 		Message: fmt.Sprintf("Hello %s from batch %s", data.Name, data.Batch),
	// 	})

	// search, isFilled := ctx.GetQuery("s")

	// if isFilled {
	// 	ctx.JSON(200, Response{
	// 		Success: true,
	// 		Message: fmt.Sprintf("Hasil pencarian dari \"%s\"", search),
	// 	})
	// } else {
	// 	ctx.JSON(404, Response{
	// 		Success: false,
	// 		Message: "Tidak item yang dicari",
	// 	})
	// }
	// })

	r.Run(":8080")
}
