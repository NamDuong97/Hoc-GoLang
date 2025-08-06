package main

import (
	"github.com/gin-gonic/gin"

	v1handler "doannam.vn/hoc-go/Group_Router/internal/api/v1/handler"

	v2handler "doannam.vn/hoc-go/Group_Router/internal/api/v2/handler"
)

func main() {
	// Tìm hiểu cách sử dụng group router trong Gin
	r := gin.Default() // tạo server mặc định

	// 1. Tạo group router tạo group api v1
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			UserHandlerV1 := v1handler.NewUserHandler()
			user.GET("/", UserHandlerV1.GetUserV1)
			user.GET("/:id", UserHandlerV1.GetUserByIdV1)
			user.POST("/", UserHandlerV1.PostUserV1)
			user.PUT("/:id", UserHandlerV1.PutUserV1)
			user.DELETE("/:id", UserHandlerV1.DeleteUserV1)
		}

		product := v1.Group("/product")
		{
			ProductHandlerV1 := v1handler.NewProductHandler()
			product.GET("/", ProductHandlerV1.GetProductV1)
			product.GET("/:id", ProductHandlerV1.GetProductByIdV1)
			product.POST("/", ProductHandlerV1.PostProductV1)
			product.PUT("/:id", ProductHandlerV1.PutProductV1)
			product.DELETE("/:id", ProductHandlerV1.DeleteProductV1)
		}
	}

	// tạo group api v2
	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/user")
		{
			UserHandlerV2 := v2handler.NewUserHandler()
			user.GET("/", UserHandlerV2.GetUserV2)
			user.GET("/:id", UserHandlerV2.GetUserByIdV2)
			user.POST("/", UserHandlerV2.PostUserV2)
			user.PUT("/:id", UserHandlerV2.PutUserV2)
			user.DELETE("/:id", UserHandlerV2.DeleteUserV2)
		}

		product := v2.Group("/product")
		{
			ProductHandlerV2 := v2handler.NewProductHandler()
			product.GET("/", ProductHandlerV2.GetProductV2)
			product.GET("/:id", ProductHandlerV2.GetProductByIdV2)
			product.POST("/", ProductHandlerV2.PostProductV2)
			product.PUT("/:id", ProductHandlerV2.PutProductV2)
			product.DELETE("/:id", ProductHandlerV2.DeleteProductV2)
		}
	}

	r.Run(":8080")
}
