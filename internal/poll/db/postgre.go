package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/poll"
	//"github.com/uks91/polls_go/internal/user"
)

var strg poll.Storage = &pollStorage{}

type pollStorage struct {
	db *sqlx.DB
}

func (p pollStorage) GetPollsList() ([]poll.Poll, error) {
	var polls []poll.Poll
	err := p.db.Select(&polls, `SELECT * from "polls"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all polls: %v", err)
	}

	return polls, nil
}

func (p pollStorage) GetPoll(id string) (poll.Poll, error) {
	//TODO implement me
	panic("implement me")
}

func (p pollStorage) GetQuestion(id string) (poll.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (p pollStorage) CreatePoll(poll *poll.Poll) (int64, error) {
	//TODO implement me
	panic("implement me")
}

/*
func (u *pollStorage) GetUserByLogin(login string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf(`SELECT "id", "login", "password", "role" FROM users WHERE "login" = '%s'`, login)
	err := u.db.Get(&usr, query)
	if err != nil {
		return usr, fmt.Errorf("user with login=%s not found: %v", login, err)
	}
	return usr, nil
}

func (u *pollStorage) GetOne(uuid string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf(`SELECT "id", "login", "role", "password" from "users" WHERE "id"=%s`, uuid)
	err := u.db.Get(&usr, query)
	if err != nil {
		fmt.Printf("failed query id: %s", query)
		return usr, fmt.Errorf("user with id=%s not found: %v", uuid, err)
	}
	fmt.Println("usr == ", usr)
	return usr, nil
}

func (u *pollStorage) GetAll() ([]user.User, error) {
	var users []user.User
	err := u.db.Select(&users, `SELECT "id", "login", "role", "password" from "users"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all users: %v", err)
	}

	return users, nil
}

func (u *pollStorage) Create(usr user.User) (user.User, error) {
	out := usr
	query := fmt.Sprintf(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES ('%s', '%s', '%s', '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr.Username, usr.PasswordHash, usr.Role)
	fmt.Println(query)
	//row := u.db.QueryRowx(query)
	exec, err := u.db.NamedExec(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES (:login, :password, :role, '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr)
	if err != nil {
		return user.User{}, fmt.Errorf("unable to create new user: %v", err)
	}
	id, _ := exec.LastInsertId()
	fmt.Printf("inserted id is %s", string(id))
	out.ID = string(id)
	return out, nil
}

func (u *pollStorage) Delete(book *user.User) error {
	//TODO implement me
	panic("implement me")
}
*/

func NewStorage(db *sqlx.DB) poll.Storage {
	return &pollStorage{
		db: db,
	}
}
