package crypto_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	password := "123456"
	hash, _ := HashPassword(password)
	assert.Truef(t, CheckPasswordHash(password, hash), "Valid hash")
}
