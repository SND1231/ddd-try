package efuser

import (
	"ddd/v1/user"
	"ddd/v1/db"
	"ddd/v1/entity"
)

type EfUserRepository struct {}

func(efuser EfUserRepository) save(user user.User) bool {
	userData := entity.User{ID: user.GetUserId(), Name: user.GetFullName()}
	
	db := db.Connection()
	defer db.Close()
	db.Create(&userData)
	if db.NewRecord(userParam) == false {
		return true
	}

	return false
}

func(efuser EfUserRepository) getUserByCode(code user.UserCode) user.User {
	var userData entity.User

	db := db.Connection()
	defer db.Close()
	db.Find(&user, code.GetUserCode()).Scan(&userData)

	return user.User{code: userData.Code, name: userData.Name}
}

func(efuser EfUserRepository) exists(user UserCode) bool {
	var userData entity.User

	db := db.Connection()
	defer db.Close()

	db.where("code = ?", user.GetUserCode())
}

