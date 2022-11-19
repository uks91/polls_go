package user

import (
	"github.com/gin-gonic/gin"
	"github.com/uks91/polls_go/internal"
)

type userHandler struct {
	service Service
}

func NewUserHandler(service Service) internal.Handler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) Register(group *gin.RouterGroup) {
	group.GET("/", h.GetUsers)
	group.GET("/:id", h.GetUser)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	h.service.GetAll()
}

func (h *userHandler) GetUser(c *gin.Context) {

}
