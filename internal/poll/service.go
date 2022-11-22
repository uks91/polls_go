package poll

import "fmt"

//type UserService interface {
//	GetAllUsers(c *gin.Context) error
//}

type Service struct {
	storage Storage
}

func (u *Service) GetPollsList() ([]PollDTO, error) {
	polls, err := u.storage.GetPollsList()
	if err != nil {
		return nil, fmt.Errorf("unable to get polls list: %v", err)
	}
	pollsdto := make([]PollDTO, len(polls))
	for i, poll := range polls {
		pollsdto[i].Id = poll.Id
		pollsdto[i].Name = poll.Name
		pollsdto[i].Description = poll.Description
	}
	return pollsdto, nil
}

func (u *Service) GetPoll(id string) (Poll, error) {
	u.storage.GetQuestions(id)
	return u.storage.GetPoll(id)
}

func (u *Service) CreatePoll(p *PollDTO) error {
	return nil
}

func NewPollService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
