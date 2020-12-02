package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"ddd/v1/user"
)

func TestCreateUser(t *testing.T) {
	user1, err := user.CreateUser("1111", "John", "Ader")

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	assert.Equal(t, "1111", user1.GetUserId(), "The two words should be the same.")
	assert.Equal(t, "John Ader", user1.GetFullName(), "The two words should be the same.")

	err = user1.UpdateUser("Bob", "Gary")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	assert.Equal(t, "Bob Gary", user1.GetFullName(), "The two words should be the same.")
}
