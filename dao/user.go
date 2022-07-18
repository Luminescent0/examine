package dao

import (
	"examine/model"
	"fmt"
)

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := dB.Table("user").Where("username=?", username).Find(&user)
	fmt.Println(err.Error)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func InsertUser(user model.User) error {
	err := dB.Table("user").Select("username", "password").Create(&user)
	if err != nil {
		fmt.Println(err.Error)
		return err.Error
	}
	return nil

}
