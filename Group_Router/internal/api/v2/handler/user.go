package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v2!"})
}

func (u *UserHandler) GetUserByIdV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v2!"})
}

func (u *UserHandler) PostUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v2!"})
}

func (u *UserHandler) PutUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v2!"})
}

func (u *UserHandler) DeleteUserV2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v2!"})
}
