package model

import (
	"project_golang/services/user/typeuser"
)

type UserModel struct {}

func (mm *UserModel)FindUser(mobile string) (*typeuser.User, error) {
	user := typeuser.User{
		"alexluan",
		"18817322111",
		"2222@qq.com",
		12,
		"shanghai",
		"http://xxx.jpg",
	}
	return &user, nil
}