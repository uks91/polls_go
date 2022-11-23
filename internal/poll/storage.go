package poll

type Storage interface {
	GetPollsList() ([]Poll, error)
	GetPoll(id string) (Poll, error)
	GetQuestion(id string) (Question, error)
	CreatePoll(poll *Poll) (string, error)
	GetQuestions(pollId string) ([]Question, error)
}
