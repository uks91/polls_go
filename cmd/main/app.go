package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/composite"
)

func main() {
	router := gin.Default()
	conn, err := sqlx.Connect("pgx", "postgresql://postgres:pgadmin@localhost:5432/polls")
	if err != nil {
		panic(err)
	}

	userComposite := composite.NewUserComposite(conn)
	userComposite.Register(router.Group("/api/user"))

	pollComposite := composite.NewPollComposite(conn)
	pollComposite.Register(router.Group("/api/polls"))

	router.Run("127.0.0.1:1001")
}
