package user

type UserDTO struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID           string `json:"id" bson:"_id,omitempty" db:"id"`
	Username     string `json:"login" bson:"username" db:"login"` // bson - for mongo
	PasswordHash string `json:"-" bson:"password" db:"password"`
	Role         string `json:"-" bson:"password" db:"role"`
	CreatedAt    string `db:"createdAt"`
	UpdatedAt    string `db:"updatedAt"`
}
