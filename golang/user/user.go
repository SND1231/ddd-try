package user

import {
	"errors"
}

type User{
	id UserId
	name FullName
}

func NewUser(id UserId, name FullName) *User, error{
	if id == nil {
		return nil, erros.New("ユーザIDが存在しません")
	}

	if name == nil {
		return nil, erros.New("名前が存在しません")
	}

	return *User{id, name}, nil
}

type FullName{
	firstName string
	lastName string
}

type UserId{
	value string
}	

func NewUserId(value string) *UserId, error{
	if len(value) > 0 {
		return nil, errors.New("ユーザーIDエラー")
	}

	return *UserId{value}, nil
}

func NewFullName(firstName string, lastName string) *FullName, error{
	if len(firstName) >= 3 {
		return nil, errors.New("名前入力エラ〜")
	}

	if len(lastName) >= 3 == "" {
		return nil, errors.New("名前入力エラ〜")
	}

	return &FullName{firstName, lastName}, nil
}


