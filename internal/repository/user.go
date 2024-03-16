package repository

import (
	"database/sql"

	"github.com/albiosz/honeycombs"
	"github.com/albiosz/honeycombs/internal/config/errs"
	"github.com/albiosz/honeycombs/internal/database"
)

var _ honeycombs.UserRepository = (*User)(nil)

type User struct {
	db *database.DB
}

func NewUser(db *database.DB) *User {
	return &User{db: db}
}

func (u *User) FindByEmail(email string) (*honeycombs.User, error) {
	user := &honeycombs.User{}
	row := u.db.SqlDB.QueryRow("SELECT email, password, nickname FROM users WHERE email = ?", email)
	if err := row.Scan(&user.Email, &user.Password, &user.Nickname); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrResourceNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u *User) Find() ([]*honeycombs.User, error) {
	var users []*honeycombs.User

	rows, err := u.db.SqlDB.Query("SELECT email, password, nickname FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user honeycombs.User
		if err := rows.Scan(&user.Email, &user.Password, &user.Nickname); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) Create(user *honeycombs.User) (*honeycombs.User, error) {
	result, err := u.db.SqlDB.Exec("INSERT INTO users (email, password, nickname) VALUES (?, ?, ?)", user.Email, user.Password, user.Nickname)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected != 1 {
		return nil, err
	}

	createdUser, err := u.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *User) Update(email string, updates honeycombs.UserUpdate) (*honeycombs.User, error) {
	user, err := u.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if updates.Nickname != nil {
		user.Nickname = *updates.Nickname
	}
	if updates.Password != nil {
		user.Password = *updates.Password
	}

	if _, err := u.db.SqlDB.Exec(`UPDATE users
		SET password = ?, nickname = ?
		WHERE email = ?
		`, user.Password, user.Nickname, user.Email,
	); err != nil {
		return nil, err
	}

	updateUser, err := u.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (u *User) Delete(email string) error {
	_, err := u.db.SqlDB.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		return err
	}
	return nil
}
