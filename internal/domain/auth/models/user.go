package models

import (
	"encoding/base64"
	"net/mail"
	"regexp"
)

type User struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (u *User) ValidateLogin() bool {
	return u.Login != ""
}

func (u *User) ValidateEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u *User) ValidatePassword() bool {
	return u.Password != ""
}

func (u *User) ValidatePhone() bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(u.Phone) && u.Phone != ""
}

func (u *User) EncodePassword() string {
	return base64.StdEncoding.EncodeToString([]byte(u.Password))
}

func (u *User) DecodePassword() (string, error) {
	s, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
