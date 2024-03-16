package honeycombs

type User struct {
	// ID int - it would be probably a better PK, because it would allow to change the Email
	Email    string
	Password string
	Nickname string
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Find() ([]*User, error)
	Create(user *User) (*User, error)
	Update(email string, updates UserUpdate) (*User, error)
	Delete(email string) error
}

type UserUpdate struct {
	Password *string
	Nickname *string
}
