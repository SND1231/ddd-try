package user

import (
	"errors"
)

type User struct{
	code UserCode
	name FullName
}

func CreateUser(code string, firstName string, lastName string) (User, error){
	var userCode UserCode
	var fullName FullName
	var err error

	userCode, err = newUserId(id)
	if err != nil {
		return User{}, err
	}

	fullName, err = newFullName(firstName, lastName)
	if err != nil {
		return User{}, err
	}

	return newUser(userCode, fullName)
}

func(u *User) UpdateUser(firstName string, lastName string) error {
	var fullName FullName
	var err error

	fullName, err = newFullName(firstName, lastName)
	if err != nil {
		return err
	}

	u.name = fullName
	return nil
}

func newUser(code userCode, name FullName) (User, error) {
	return User{code, name}, nil
}

func(u User) GetUserCode() string {
	return u.code.value
}

func(u User) GetFullName() string {
	return u.name.firstName + " " + u.name.lastName
}

type FullName struct{
	firstName string
	lastName string
}

func newFullName(firstName string, lastName string) (FullName, error) {
	if len(firstName) < 3 {
		return FullName{}, errors.New("名前入力エラ〜")
	}

	if len(lastName) < 3 {
		return FullName{}, errors.New("名前入力エラ〜")
	}

	return FullName{firstName, lastName}, nil
}


type UserCode struct{
	value string
}	

func newUserCode(value string) (UserCode, error) {
	if len(value) == 0 {
		return UserId{}, errors.New("ユーザーIDエラー")
	}

	return UserCode{value}, nil
}

type IUserRepository interface {
	save(User) bool
	getUserByCode(UserCode) User
	exists(User) bool
}

