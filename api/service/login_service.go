package service

import (
	"golang.org/x/crypto/bcrypt"
)

func (a *App) Login(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}