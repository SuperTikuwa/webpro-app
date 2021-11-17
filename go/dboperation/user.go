package dboperation

import "github.com/SuperTikuwa/webpro-app/model"

func CreateUser(user *model.User) error {
	db := gormConnect()
	defer db.Close()

	return db.Create(*user).Error
}

func Login(user *model.User) error {
	db := gormConnect()
	defer db.Close()

	return db.Where("name=? and password=?", user.Name, user.Password).First(user).Error
}

func SelectUserByName(name string) (model.User, error) {
	db := gormConnect()
	defer db.Close()

	var user model.User
	err := db.Where("name=?", name).First(&user).Error
	return user, err
}
