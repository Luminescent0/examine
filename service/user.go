package service

import (
	"database/sql"
	"errors"
	"examine/dao"
	"examine/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		fmt.Println(err)
		flag := errors.Is(err, gorm.ErrRecordNotFound)
		if !flag {
			return false, err
		}
		fmt.Println(username) //验证是否ErrNoRows
		return false, err
	}

	flag := ComparePassword(user.Password, []byte(password))
	if !flag {
		return false, nil
	}
	fmt.Println("验证密码成功")
	return true, nil
}

func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func Register(user model.User) error {
	password, err := Cipher(user)
	if err != nil {
		return err
	}
	user.Password = password
	err = dao.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

// Bcrypt 密码加盐
type Bcrypt struct {
	cost int
}

func Cipher(user model.User) (string, error) {
	password := []byte(user.Password)
	nowG := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	fmt.Println("加密后", string(hashedPassword), "耗时", time.Now().Sub(nowG))

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UsernameIsExist(username string) error {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			fmt.Println("不存在")
			return err
		}
		return err
	}
	return nil
}
