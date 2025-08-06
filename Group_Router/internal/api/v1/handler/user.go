package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all user v1!"})
}

func (u *UserHandler) GetUserByIdV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user by id v1!"})
}

func (u *UserHandler) PostUserV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Post user v1!"})
}

func (u *UserHandler) PutUserV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Put user v1!"})
}

func (u *UserHandler) DeleteUserV1(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{"message": "Delete user v1!"})
}
