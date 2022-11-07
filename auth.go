package main

import (
	"errors"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"nick": "!Aa123456",
	}
}

func checkUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	}

	return errors.New("password is not correct")
}

func Auth(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		return CheckPassword(UserData[username], password)
	}

	return errors.New("user is not exist")
}
