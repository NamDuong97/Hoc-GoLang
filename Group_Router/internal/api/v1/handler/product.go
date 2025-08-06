package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v1!"})
}

func (u *ProductHandler) GetProductByIdV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get usProducter by id v1!"})
}

func (u *ProductHandler) PostProductV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Post Product v1!"})
}

func (u *ProductHandler) PutProductV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Put Product v1!"})
}

func (u *ProductHandler) DeleteProductV1(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "Delete Product v1!"})
}
