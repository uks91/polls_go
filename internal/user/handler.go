package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uks91/polls_go/internal"
	"net/http"
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
	group.POST("/login", h.LogIn)
	group.POST("/registration", h.SignIn)

}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUser(c *gin.Context) {
	id, b := c.Params.Get("id")
	if !b {
		// TODO: log
		return
	}
	// TODO: check if id is int
	usr, err := h.service.GetOne(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	fmt.Println(usr)
	c.JSON(http.StatusOK, usr)
}

func (h *userHandler) LogIn(c *gin.Context) {
	var usr UserDTO
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.String(http.StatusTeapot, "I'm a teapot")
		fmt.Printf("I'm a teapot: %v", err)
		return
	}
	//h.service
	err := h.service.LogIn(&usr)
	if err != nil {
		c.String(http.StatusMethodNotAllowed, "login failed")
		return
	}
	c.String(http.StatusOK, "It's ok!")
}

func (h *userHandler) SignIn(c *gin.Context) {
	var usr UserDTO
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.String(http.StatusTeapot, "I'm a teapot")
		fmt.Printf("I'm a teapot: %v", err)
		return
	}
	//h.service
	err := h.service.SignIn(&usr)
	if err != nil {
		c.String(http.StatusMethodNotAllowed, "registration failed")
		return
	}
	c.String(http.StatusOK, "It's ok!")
}
