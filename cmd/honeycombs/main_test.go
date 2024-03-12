package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVariables(t *testing.T) {
	main()

	assert.NotEmpty(t, os.Getenv("DB_USERNAME"))
	assert.NotEmpty(t, os.Getenv("DB_PORT"))
	assert.NotEmpty(t, os.Getenv("DB_NAME"))
	assert.NotEmpty(t, os.Getenv("DB_USERNAME"))
	assert.NotEmpty(t, os.Getenv("DB_PASSWORD"))
}
