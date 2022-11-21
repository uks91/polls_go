package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/user"
)

type userStorage struct {
	db *sqlx.DB
}

func (u *userStorage) GetUserByLogin(login string) (user.User, error) {
	var usr user.User
	err := u.db.Get(&usr, fmt.Sprintf(`SELECT "id", "login", "role", "password" from "users" WHERE "login"=%s`, login))
	if err != nil {
		return usr, fmt.Errorf("user with login=%s not found: %v", login, err)
	}
	//fmt.Println("usr == ", usr)
	return usr, nil
}

func (u *userStorage) GetOne(uuid string) (user.User, error) {
	var usr user.User
	err := u.db.Get(&usr, fmt.Sprintf(`SELECT "id", "login", "role", "password" from "users" WHERE "id"=%s`, uuid))
	if err != nil {
		return usr, fmt.Errorf("user with id=%s not found: %v", uuid, err)
	}
	fmt.Println("usr == ", usr)
	return usr, nil
}

func (u *userStorage) GetAll() ([]user.User, error) {
	var users []user.User
	err := u.db.Select(&users, `SELECT "id", "login", "role", "password" from "users"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all users: %v", err)
	}

	return users, nil
}

func (u *userStorage) Create(book *user.UserDTO) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userStorage) Delete(book *user.User) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(db *sqlx.DB) user.Storage {
	return &userStorage{
		db: db,
	}
}
