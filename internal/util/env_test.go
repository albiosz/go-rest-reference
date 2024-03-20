package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupEnvVariables(t *testing.T) {

	t.Run("file exists", func(t *testing.T) {
		projectRoot := ProjectRoot()

		err := SetupEnvVariables(projectRoot + "/.env")

		assert.NoError(t, err)
		assert.NotEmpty(t, os.Getenv("DB_USERNAME"))
		assert.NotEmpty(t, os.Getenv("DB_PORT"))
		assert.NotEmpty(t, os.Getenv("DB_NAME"))
		assert.NotEmpty(t, os.Getenv("DB_USERNAME"))
		assert.NotEmpty(t, os.Getenv("DB_PASSWORD"))
	})

	t.Run("file does not exist", func(t *testing.T) {
		err := SetupEnvVariables("nonexistent.env")

		assert.Error(t, err)
	})
}
