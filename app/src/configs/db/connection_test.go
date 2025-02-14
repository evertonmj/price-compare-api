package configs_db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	client := NewConnection()
	assert.NotNil(t, client, "Expected non-nil Redis client")
	assert.Equal(t, "127.0.0.1:6379", client.Options().Addr, "Expected address to be 127.0.0.1:6379")
	assert.Equal(t, "default", client.Options().Username, "Expected username to be default")
	assert.Equal(t, "secret", client.Options().Password, "Expected password to be secret")
}
