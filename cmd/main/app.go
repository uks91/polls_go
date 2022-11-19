package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uks91/polls_go/internal/user"
	"github.com/uks91/polls_go/internal/user/db"
)

func main() {
	router := gin.Default()
	//hand := handlers.NewHandler()
	//r := hand.InitRoutes()
	//r.Run("127.0.0.1:1001")

	userStorage := db.NewStorage()
	userService := user.NewUserService(userStorage)
	userHandler := user.NewUserHandler(userService)
	userHandler.Register(router.Group("/api/user"))

	router.Run("127.0.0.1:1001")
}
