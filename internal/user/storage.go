package user

type Storage interface {
	GetOne(uuid string) (User, error)
	GetAll() ([]User, error)
	Create(usr *UserDTO) (User, error)
	GetUserByLogin(login string) (User, error)
	Delete(usr *User) error
}
