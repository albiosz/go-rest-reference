package database

import (
	"testing"

	"github.com/albiosz/honeycombs/internal/util"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *DB {
	util.SetupEnvVariables(util.ProjectRoot() + "/.env")
	return Get()
}

func TestGet(t *testing.T) {
	db := setupTestDB()
	defer db.SqlDB.Close()

	t.Run("env variables are set", func(t *testing.T) {
		pingErr := db.SqlDB.Ping()
		assert.Nil(t, pingErr)
	})
}

func TestClear(t *testing.T) {
	db := setupTestDB()
	defer db.SqlDB.Close()

	countRows := func(db *DB) int {
		var numOfRows int
		row := db.SqlDB.QueryRow("SELECT COUNT(*) FROM users;")
		err := row.Scan(&numOfRows)
		assert.NoError(t, err)
		return numOfRows
	}

	t.Run("clear Database", func(t *testing.T) {
		db.SqlDB.Exec("INSERT INTO users (email, password, nickname) VALUES ('new@email.de', 'pass', 'nick')")

		assert.Greater(t, countRows(db), 0)

		db.Clear()

		assert.Equal(t, 0, countRows(db))
	})
}
