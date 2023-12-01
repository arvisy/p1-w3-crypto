package handler

import (
	"cryptoss/config"
	"cryptoss/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users(username, password) VALUES(?, ?)", user.Username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
