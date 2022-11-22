package composite

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal"
	"github.com/uks91/polls_go/internal/user"
	"github.com/uks91/polls_go/internal/user/db"
)

type UserComposite struct {
	Storage user.Storage
	Service user.Service
	Handler internal.Handler
}

func NewUserComposite(database *sqlx.DB) *UserComposite {
	userStorage := db.NewStorage(database)
	userService := user.NewUserService(userStorage)
	userHandler := user.NewUserHandler(userService)
	//userHandler.Register(group)
	return &UserComposite{
		Storage: userStorage,
		Service: userService,
		Handler: userHandler,
	}
}

func (c *UserComposite) Register(group *gin.RouterGroup) {
	c.Handler.Register(group)
}
