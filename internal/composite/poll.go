package composite

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal"
	"github.com/uks91/polls_go/internal/poll"
	"github.com/uks91/polls_go/internal/poll/db"
)

type PollComposite struct {
	Storage poll.Storage
	Service *poll.Service
	Handler internal.Handler
}

func NewPollComposite(database *sqlx.DB) *PollComposite {
	pollStorage := db.NewStorage(database)
	pollService := poll.NewPollService(pollStorage)
	pollHandler := poll.NewPollHandler(pollService)
	return &PollComposite{
		Storage: pollStorage,
		Service: pollService,
		Handler: pollHandler,
	}
}

func (c *PollComposite) Register(group *gin.RouterGroup) {
	c.Handler.Register(group)
}
