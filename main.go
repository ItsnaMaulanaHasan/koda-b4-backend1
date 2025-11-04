package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Talent struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Batch   int    `json:"batch" binding:"required"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type ResponseGetAllTalents struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Talent `json:"data"`
}

type ResponseGetTalentById struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Talent `json:"data"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// type Query struct {
// 	Name  string `form:"name" binding:"len=10"`
// 	Batch string `form:"batch"`
// }

var talents = []Talent{
	{
		Id:      1,
		Name:    "Itsna",
		Batch:   4,
		Phone:   "12345678",
		Address: "Pati",
	},
	{
		Id:      2,
		Name:    "Federus",
		Batch:   4,
		Phone:   "12345678",
		Address: "Aceh",
	},
	{
		Id:      3,
		Name:    "Ari",
		Batch:   4,
		Phone:   "12345678",
		Address: "Depok",
	},
	{
		Id:      4,
		Name:    "Yoga",
		Batch:   4,
		Phone:   "12345678",
		Address: "Cibubur",
	},
	{
		Id:      5,
		Name:    "Fiki",
		Batch:   4,
		Phone:   "12345678",
		Address: "Sidoarjo",
	},
	{
		Id:      6,
		Name:    "Sidik",
		Batch:   4,
		Phone:   "12345678",
		Address: "Jepara",
	},
	{
		Id:      7,
		Name:    "Anggi",
		Batch:   4,
		Phone:   "12345678",
		Address: "Jambi",
	},
}

func main() {
	r := gin.Default()

	// Mendapatkan semua data talent
	r.GET("/talents", func(ctx *gin.Context) {
		ctx.JSON(200, ResponseGetAllTalents{
			Success: true,
			Message: "Success get data talents",
			Data:    talents,
		})
	})

	// Mendapatkan data talent berdasarkan id
	r.GET("/talents/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		var result Talent

		found := false
		for i := range talents {
			if talents[i].Id == idInt {
				result = talents[i]
				found = true
			}
		}

		if found {
			ctx.JSON(200, ResponseGetTalentById{
				Success: true,
				Message: "Talent data found",
				Data:    result,
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: "Talent data not found",
			})
		}
	})

	// Menambah user baru
	r.POST("/talents", func(ctx *gin.Context) {
		var body Talent

		err := ctx.BindJSON(&body)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("Failed to add talent: %v", err.Error()),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"success": true,
			"message": "Success add data talent",
		})

		talents = append(talents, body)

		fmt.Println("Hasil tambah data talent: ")
		fmt.Println(talents)
	})

	// Mengedit data talent berdasarkan id
	r.PATCH("/talents/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		var body Talent

		err := ctx.BindJSON(&body)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("Failed to add talent: %v", err.Error()),
			})
			return
		}

		found := false
		for i := range talents {
			if talents[i].Id == idInt {
				talents[i].Name = body.Name
				talents[i].Batch = body.Batch
				talents[i].Phone = body.Phone
				talents[i].Address = body.Address
				found = true
			}
		}

		if found {
			ctx.JSON(200, Response{
				Success: true,
				Message: "Talent data successfully updated",
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: "Talent data not found",
			})
		}

		fmt.Println("Hasil edit data talent: ")
		fmt.Println(talents)
	})

	r.DELETE("/talents/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idInt, _ := strconv.Atoi(id)

		found := false
		for i := range talents {
			if talents[i].Id == idInt {
				talents = append(talents[:i], talents[i+1:]...)
				found = true
				break
			}
		}

		if found {
			ctx.JSON(200, Response{
				Success: true,
				Message: "talent data successfully deleted",
			})
		} else {
			ctx.JSON(404, Response{
				Success: false,
				Message: "Talent data not found",
			})
		}

		fmt.Println("Hasil delete data: ")
		fmt.Println(talents)
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
