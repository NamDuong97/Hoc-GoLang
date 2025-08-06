package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v1!"})
}
func GetUserByIdV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user by id v1!"})
}

func main() {
	r := gin.Default() // tạo server mặc định

	// 1. Tìm hiểu các phương thức HTTP cơ bản và cách định nghĩa route trong Gin
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// param thì không thể truyền thiếu ví dụ: localhost:8080/ => sẽ báo lỗi mà phải truyền đủ param
	// ví dụ: localhost:8080/user/123 => sẽ trả về id là
	r.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{"message": "nguoi dung co id: " + id})
	})

	// query thì có thể truyền thiếu, ví dụ: localhost:8080/product?price=100&color=red
	// nếu truyền thiếu thì sẽ trả về giá trị rỗng cho param nào không truyền
	r.GET("/product/:prduct_name", func(ctx *gin.Context) {
		productName := ctx.Param("prduct_name")
		price := ctx.Query("price")
		color := ctx.Query("color")
		ctx.JSON(200, gin.H{
			"message": "san pham co ten: " + productName,
			"price":   price,
			"color":   color,
		})
	})

	r.Run(":8080")
}
