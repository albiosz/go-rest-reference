package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectRoot(t *testing.T) {
	got := ProjectRoot()
	projectName := "honeycombs"
	pathEndsWith := got[len(got)-len(projectName):]

	assert.Equal(t, pathEndsWith, projectName)
}
