package handlers

import "github.com/gin-gonic/gin"

// Структуру лучше сделать как в learn_go
type Handler interface {
	Register(group *gin.RouterGroup)
}
