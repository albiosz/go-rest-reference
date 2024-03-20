package service

import "github.com/albiosz/honeycombs"

var _ honeycombs.UserService = (*User)(nil)

type User struct {
	userRepo honeycombs.UserRepository
}

func NewUser(userRepo honeycombs.UserRepository) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) FindByID(id uint) (*honeycombs.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *User) Find() ([]*honeycombs.User, error) {
	return u.userRepo.Find()
}

func (u *User) Create(user *honeycombs.User) (*honeycombs.User, error) {
	return u.userRepo.Create(user)
}

func (u *User) Update(id uint, updates honeycombs.UserUpdate) (*honeycombs.User, error) {
	return u.userRepo.Update(id, updates)
}

func (u *User) Delete(id uint) error {
	return u.userRepo.Delete(id)
}
