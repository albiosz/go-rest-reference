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
		user, err := userRepo.FindByID(1)

		assert.NoError(t, err)
		assert.Equal(t, "user1@mail.com", user.Email)
		assert.Equal(t, "user1", user.Nickname)
		assert.Equal(t, "password1", user.Password)
	})

	t.Run("user does not exist", func(t *testing.T) {
		user, err := userRepo.FindByID(99999)

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
		assert.NotEmpty(t, createdUser.ID)
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
		incorrectID := uint(9999)

		updatedUser, err := userRepo.Update(incorrectID, userUpdate)
		assert.Error(t, err)
		assert.Nil(t, updatedUser)
	})

	t.Run("success", func(t *testing.T) {
		correctID := uint(1)

		updatedUser, err := userRepo.Update(correctID, userUpdate)
		assert.NoError(t, err)
		assert.Equal(t, *userUpdate.Password, updatedUser.Password)
	})
}

func TestDelete(t *testing.T) {
	db := test.SetupDB()
	defer db.SqlDB.Close()
	userRepo := NewUser(db)

	t.Run("fail - email does not exist", func(t *testing.T) {
		const incorrectID = 9999
		err := userRepo.Delete(incorrectID)
		assert.NoError(t, err)
	})

	t.Run("success", func(t *testing.T) {
		const correctID = 2

		userBeforeDelete, err := userRepo.FindByID(correctID)
		assert.NoError(t, err)
		assert.NotEmpty(t, userBeforeDelete)

		err = userRepo.Delete(correctID)
		assert.NoError(t, err)

		userAfterDelete, err := userRepo.FindByID(correctID)
		assert.Error(t, err)
		assert.Nil(t, userAfterDelete)
	})
}
