package poll

type Storage interface {
	GetPollsList() ([]Poll, error)
	GetPoll(id string) (Poll, error)
	GetQuestion(id string) (Question, error)
	CreatePoll(poll *Poll) (int64, error)
}
