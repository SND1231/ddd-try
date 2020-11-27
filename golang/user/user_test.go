package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"ddd/v1/user"
)

func TestNewUser(t *testing.T) {
	userId, _ := user.NewUserId("1111")
	fullName, _ := user.NewFullName("John", "Ader")

	assert.Equal(t, "1111", userId.value, "The two words should be the same.")

	user1, _ := user.NewUser(userId, fullName)

	assert.Equal(t, "1111", user1, "The two words should be the same.")
	
}
