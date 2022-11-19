package db

import "github.com/uks91/polls_go/internal/user"

type userStorage struct {
	// db ...
}

func (u *userStorage) GetOne(uuid string) *user.User {
	//TODO implement me
	panic("implement me")
}

func (u *userStorage) GetAll(limit, offset int) []*user.User {
	//TODO implement me
	panic("implement me")
}

func (u *userStorage) Create(book *user.User) *user.User {
	//TODO implement me
	panic("implement me")
}

func (u *userStorage) Delete(book *user.User) error {
	//TODO implement me
	panic("implement me")
}

// first argument - db
func NewStorage() user.Storage {
	return &userStorage{
		// db: db
	}
}
