package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v2!"})
}

func (u *ProductHandler) GetProductByIdV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v2!"})
}

func (u *ProductHandler) PostProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v2!"})
}

func (u *ProductHandler) PutProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v2!"})
}

func (u *ProductHandler) DeleteProductV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all Product v2!"})
}
