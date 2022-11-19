package user

//type UserService interface {
//	GetAllUsers(c *gin.Context) error
//}

type userService struct {
	storage Storage
}

func (u *userService) GetAll() []*User {
	u.storage.GetAll(1, 1)
	return nil
}

func (u *userService) GetOne(id string) User {
	//TODO implement me
	panic("implement me")
}

type Service interface {
	GetAll() []*User
	GetOne(id string) User
}

func NewUserService(storage Storage) Service {
	return &userService{
		storage: storage,
	}
}
