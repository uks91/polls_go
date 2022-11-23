package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/uks91/polls_go/internal/user"
	"strconv"
)

type userStorage struct {
	db *sqlx.DB
}

func (u *userStorage) GetUserByLogin(login string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf(`SELECT "id", "login", "password", "role" FROM users WHERE "login" = '%s'`, login)
	err := u.db.Get(&usr, query)
	if err != nil {
		return usr, fmt.Errorf("user with login=%s not found: %v", login, err)
	}
	return usr, nil
}

func (u *userStorage) GetOne(uuid string) (user.User, error) {
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

func (u *userStorage) GetAll() ([]user.User, error) {
	var users []user.User
	err := u.db.Select(&users, `SELECT "id", "login", "role", "password" from "users"`)
	if err != nil {
		return nil, fmt.Errorf("unable to get all users: %v", err)
	}

	return users, nil
}

func (u *userStorage) Create(usr user.User) (user.User, error) {
	out := usr
	query := fmt.Sprintf(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES ('%s', '%s', '%s', '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr.Username, usr.PasswordHash, usr.Role)
	fmt.Println(query)
	//row := u.db.QueryRowx(query)
	exec, err := u.db.NamedExec(`INSERT INTO "users" ("login", "password", "role", "createdAt", "updatedAt")  VALUES (:login, :password, :role, '2022-11-22', '2022-11-22') RETURNING ("id", "login", "password", "role")`, usr)
	if err != nil {
		return user.User{}, fmt.Errorf("unable to create new user: %v", err)
	}
	idInt, _ := exec.LastInsertId()
	id := strconv.FormatInt(idInt, 10)
	fmt.Printf("inserted id is %s", id)
	out.ID = id
	return out, nil
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
