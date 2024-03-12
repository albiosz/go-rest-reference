package database

import (
	"testing"

	"github.com/albiosz/honeycombs/internal/config"
	"github.com/albiosz/honeycombs/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("env variables are set", func(t *testing.T) {
		config.SetupEnvVariables(util.ProjectRoot() + "/.env")
		db := Get()
		defer db.Close()

		pingErr := db.Ping()
		assert.Nil(t, pingErr)
	})
}
