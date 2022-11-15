package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uks91/polls_go/internal/handlers"
)

func main() {
	router := gin.Default()
	//hand := handlers.NewHandler()
	//r := hand.InitRoutes()
	//r.Run("127.0.0.1:1001")

	userHandler := handlers.NewUserHandler()
	userHandler.Register(router.Group("/"))

	router.Run("127.0.0.1:1001")
}
