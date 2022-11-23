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
	pollObj, err := u.storage.GetPoll(id)
	if err != nil {
		return Poll{}, err
	}
	return pollObj, nil
}

func (u *Service) CreatePoll(p *PollDTO) error {
	pollObj := Poll{
		Name:        p.Name,
		Description: p.Description,
		Question:    nil,
		CreatedAt:   "2022-11-22",
		UpdatedAt:   "2022-11-22",
	}
	questions := make([]Question, len(p.Questions))
	getOptions := func(dto *[]string) []Option {
		if dto == nil {
			return make([]Option, 0)
		}
		out := make([]Option, len(*dto))
		for i, str := range *dto {
			out[i] = Option{
				Text:      str,
				CreatedAt: pollObj.CreatedAt,
				UpdatedAt: pollObj.UpdatedAt,
			}
		}
		return out
	}
	for i, q := range p.Questions {
		questions[i] = Question{
			Text:      q.Text,
			Type:      fmt.Sprint(q.Type),
			Options:   getOptions(&q.Options),
			CreatedAt: pollObj.CreatedAt,
			UpdatedAt: pollObj.UpdatedAt,
		}
	}
	pollObj.Question = questions
	_, err := u.storage.CreatePoll(&pollObj)
	if err != nil {
		fmt.Printf("Creating poll: %v", err)
		return err
	}
	return nil
}

func NewPollService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
