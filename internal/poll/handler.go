package poll

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uks91/polls_go/internal"
	"net/http"
)

type pollHandler struct {
	service *Service
}

func NewPollHandler(service *Service) internal.Handler {
	return &pollHandler{
		service: service,
	}
}

func (h *pollHandler) Register(group *gin.RouterGroup) {
	group.GET("/", h.GetPollsList)
	group.GET("/:poll_id", h.GetPoll)
	group.POST("/new", h.CreatePoll)
}

func (h *pollHandler) GetPollsList(c *gin.Context) {
	polls, err := h.service.GetPollsList()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, polls)
}

func (h *pollHandler) GetPoll(c *gin.Context) {
	id, b := c.Params.Get("poll_id")
	if !b {
		// TODO: log
		c.String(http.StatusBadRequest, "something goes wrong")
		return
	}
	//// TODO: check if id is int
	pollObj, err := h.service.GetPoll(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	fmt.Println(pollObj)
	c.JSON(http.StatusOK, pollObj)
}

func (h *pollHandler) CreatePoll(c *gin.Context) {
	var pollObj PollDTO
	if err := c.ShouldBindJSON(&pollObj); err != nil {
		c.String(http.StatusTeapot, "I'm a teapot")
		fmt.Printf("I'm a teapot: %v", err)
		return
	}
	err := h.service.CreatePoll(&pollObj)
	if err != nil {
		c.String(http.StatusMethodNotAllowed, "unable to create a new poll")
		return
	}
	c.String(http.StatusOK, "It's ok!")
}
