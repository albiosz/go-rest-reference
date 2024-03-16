package repository

import (
	"testing"

	"github.com/albiosz/honeycombs"
	"github.com/albiosz/honeycombs/internal/config/errs"
	"github.com/albiosz/honeycombs/internal/util"
	"github.com/albiosz/honeycombs/internal/util/test"
	"github.com/stretchr/testify/assert"
)

func TestFindByEmail(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	t.Run("user exists", func(t *testing.T) {
		user, err := userRepo.FindByEmail("user1@mail.com")

		assert.NoError(t, err)
		assert.Equal(t, "user1@mail.com", user.Email)
		assert.Equal(t, "user1", user.Nickname)
		assert.Equal(t, "password1", user.Password)
	})

	t.Run("user does not exist", func(t *testing.T) {
		user, err := userRepo.FindByEmail("not-existent-email")

		assert.Error(t, err)
		assert.ErrorIs(t, err, errs.ErrResourceNotFound)
		assert.Nil(t, user)
	})
}

func TestFind(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	t.Run("Success", func(t *testing.T) {
		users, err := userRepo.Find()
		assert.NoError(t, err)
		assert.Greater(t, len(users), 1)
	})
}

func TestCreate(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	t.Run("fail - primary key (email) duplication", func(t *testing.T) {
		newUser := honeycombs.User{
			Email:    "user1@mail.com", // the email is already present in seeded db
			Password: "pass",
			Nickname: "NewGuy",
		}

		createdUser, err := userRepo.Create(&newUser)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Duplicate entry")
		assert.Nil(t, createdUser)
	})

	t.Run("success - not all dates are passed - user with empty data is inserted", func(t *testing.T) {
		newUser := honeycombs.User{}

		createdUser, err := userRepo.Create(&newUser)
		assert.NoError(t, err)
		assert.Empty(t, createdUser)
	})

	t.Run("success", func(t *testing.T) {
		newUser := honeycombs.User{
			Email:    "new@email.com",
			Password: "pass",
			Nickname: "NewGuy",
		}

		createdUser, err := userRepo.Create(&newUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdUser)
	})
}

func TestUpdate(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	userUpdate := honeycombs.UserUpdate{
		Password: util.NewPtr("new-password"),
	}

	t.Run("fail - incorrect emial", func(t *testing.T) {
		incorrectEmail := "incorrect@email.com"

		updatedUser, err := userRepo.Update(incorrectEmail, userUpdate)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
	})

	t.Run("success", func(t *testing.T) {
		email := "user1@mail.com"

		updatedUser, err := userRepo.Update(email, userUpdate)
		assert.NoError(t, err)
		assert.Equal(t, *userUpdate.Password, updatedUser.Password)
	})
}

func TestDelete(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	t.Run("fail - email does not exist", func(t *testing.T) {
		const email = "does@not.exist"
		err := userRepo.Delete(email)
		assert.NoError(t, err)
	})

	t.Run("success", func(t *testing.T) {
		const email = "user2@mail.com"

		userBeforeDelete, err := userRepo.FindByEmail(email)
		assert.NoError(t, err)
		assert.NotEmpty(t, userBeforeDelete)

		err = userRepo.Delete(email)
		assert.NoError(t, err)

		userAfterDelete, err := userRepo.FindByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, userAfterDelete)
	})
}
