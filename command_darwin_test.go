package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCommand_ExecuteStderr(t *testing.T) {
	cmd := NewCommand(">&2 echo hello")

	err := cmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, "hello\n", cmd.Stderr())
}

func TestCommand_WithTimeout(t *testing.T) {
	cmd := NewCommand("sleep 0.5;", WithTimeout(5*time.Millisecond))

	err := cmd.Execute()

	assert.NotNil(t, err)
	assert.Equal(t, "Command timed out after 5ms", err.Error())
}

func TestCommand_WithValidTimeout(t *testing.T) {
	cmd := NewCommand("sleep 0.01;", WithTimeout(500*time.Millisecond))

	err := cmd.Execute()

	assert.Nil(t, err)
}
