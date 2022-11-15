package handlers

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (h *UserHandler) Register(group *gin.RouterGroup) {
	group.GET("/", h.GetUsers)
	group.GET("/:id", h.GetUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {

}

func (h *UserHandler) GetUser(c *gin.Context) {

}