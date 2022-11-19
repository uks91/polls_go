package user

type Storage interface {
	GetOne(uuid string) *User
	GetAll(limit, offset int) []*User
	Create(book *User) *User
	Delete(book *User) error
}
