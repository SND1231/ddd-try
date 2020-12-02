package usrservice

import (
	"testing"
	"ddd/v1/user"
	"ddd/v1/usrservice"
)


func TestCheckUser(t *testing.T) {
	user1, _ := user.CreateUser("1111", "John", "Ader")

	usrservice.Users = append(usrservice.Users, user1)
	result := usrservice.CheckUser(user1)

	if !result {
		t.Fatalf("failed test")
	}
}
