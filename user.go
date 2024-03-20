package honeycombs

type User struct {
	ID       uint
	Email    string
	Password string
	Nickname string
}

type UserRepository interface {
	FindByID(id uint) (*User, error)
	Find() ([]*User, error)
	Create(user *User) (*User, error)
	Update(id uint, updates UserUpdate) (*User, error)
	Delete(id uint) error
}

type UserUpdate struct {
	Email    *string
	Password *string
	Nickname *string
}
