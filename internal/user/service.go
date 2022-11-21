package user

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
	//usr, err := u.storage.GetOne(id)
	//return usr, err
	login, err := u.storage.GetUserByLogin(usr.Username)
	if err != nil {
		return err
	}
	
	return nil
}

func (u *userService) SignIn(usr *UserDTO) error {
	//usr, err := u.storage.GetOne(id)
	//return usr, err
	return nil
}

func NewUserService(storage Storage) Service {
	return &userService{
		storage: storage,
	}
}
