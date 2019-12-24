package logic

import "project_golang/services/user/model"

type UserLogic struct {
	UserModel model.UserModel
}


func (ll *UserLogic)GetUser(mobile string) (interface{}, error) {
	user, _ := ll.UserModel.FindUser(mobile)
	return user, nil
}