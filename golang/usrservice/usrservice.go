package usrservice

import (
	"ddd/v1/user"
	"fmt"
)

var Users []user.User

func CheckUser(user user.User) bool {
	count := len(Users)
	fmt.Println(count)
	for i :=0; i < count; i++ {
		if user.GetUserId() == Users[i].GetUserId(){
			return true
		}
	}
	return false
}
