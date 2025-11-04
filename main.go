package main

import (
	"github.com/gin-gonic/gin"
)

type Talent struct {
	Id      int
	Name    string
	Batch   int
	Phone   string
	Address string
}

type Response struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Talent `json:"data"`
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
		ctx.JSON(200, Response{
			Success: true,
			Message: "Success get data talents",
			Data:    talents,
		})
	})

	// Mendapatkan data talent berdasarkan id

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
