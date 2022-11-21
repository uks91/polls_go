package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/user"
	"github.com/uks91/polls_go/internal/user/db"
)

func main() {
	router := gin.Default()
	conn, err := sqlx.Connect("pgx", "postgresql://postgres:pgadmin@localhost:5432/polls")
	if err != nil {
		panic(err)
	}

	userStorage := db.NewStorage(conn)
	userService := user.NewUserService(userStorage)
	userHandler := user.NewUserHandler(userService)
	userHandler.Register(router.Group("/api/user"))

	router.Run("127.0.0.1:1001")
}
