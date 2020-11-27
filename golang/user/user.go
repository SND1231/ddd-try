package user

import (
	"errors"
)

type User struct{
	id UserId
	name FullName
}

func NewUser(id UserId, name FullName) (User, error) {
	return User{id, name}, nil
}

type FullName struct{
	firstName string
	lastName string
}

func NewFullName(firstName string, lastName string) (FullName, error) {
	if len(firstName) >= 3 {
		return FullName{}, errors.New("名前入力エラ〜")
	}

	if len(lastName) >= 3 {
		return FullName{}, errors.New("名前入力エラ〜")
	}

	return FullName{firstName, lastName}, nil
}


type UserId struct{
	value string
}	

func NewUserId(value string) (UserId, error) {
	if len(value) > 0 {
		return UserId{}, errors.New("ユーザーIDエラー")
	}

	return UserId{value}, nil
}



