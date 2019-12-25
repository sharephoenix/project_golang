package logic

import (
	"project_golang/services/user/model"
)

type UserLogic struct {
	UserModel model.UserModel
}


func (ll *UserLogic)GetUser(mobile string) (interface{}, error) {
	user, err := ll.UserModel.FindUser(mobile)
	return user, err
}

func (ll *UserLogic)Register(mobile string) (interface{}, error) {
	user, err := ll.UserModel.Register(mobile)
	return user, err
}