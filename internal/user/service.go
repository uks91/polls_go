package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//type UserService interface {
//	GetAllUsers(c *gin.Context) error
//}

type Service interface {
	GetAll() ([]User, error)
	GetOne(id string) (User, error)
	LogIn(usr *UserDTO) error
	SignIn(usr *UserDTO) error
}

type userService struct {
	storage Storage
}

func (u *userService) GetAll() ([]User, error) {
	users, err := u.storage.GetAll()
	return users, err
}

func (u *userService) GetOne(id string) (User, error) {
	usr, err := u.storage.GetOne(id)
	return usr, err
}

func (u *userService) LogIn(usr *UserDTO) error {
	login, err := u.storage.GetUserByLogin(usr.Username)
	if err != nil {
		fmt.Printf("error during logging in: %v", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(usr.Password))
	if err != nil {
		return fmt.Errorf("password is incorrect")
	}

	return nil
}

func (u *userService) SignIn(usr *UserDTO) error {
	_, err := u.storage.GetUserByLogin(usr.Username)
	if err == nil {
		return fmt.Errorf("user %s is already exists", usr.Username)
	}
	role := "USER"
	if usr.Username == "admin" {
		role = "ADMIN"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 5)
	if err != nil {
		return fmt.Errorf("unable to create hash: %v", err)
	}
	newUser := User{
		Username:     usr.Username,
		PasswordHash: string(hash),
		Role:         role,
	}
	create, err := u.storage.Create(newUser)
	if err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}
	fmt.Printf("created user with id = ", create.ID)
	return nil
}

func NewUserService(storage Storage) Service {
	return &userService{
		storage: storage,
	}
}
