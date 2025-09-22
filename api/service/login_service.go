package service

import (
	"Bell/api/repository"
)

func (a *App) Login(pass string) bool {
	setting, _ := repository.QuerySetting(a.Db)
	res := repository.CheckPasswordHash(setting.Password, pass)
	return res
}