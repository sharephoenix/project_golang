package logic

import (
	"project_golang/services/user/model"
	"project_golang/services/user/typeuser"
)

type UserLogic struct {
	UserModel model.UserModel
}


func (ll *UserLogic)GetUser(mobile string) (interface{}, error) {
	user, err := ll.UserModel.FindUser(mobile)
	return user, err
}

func (ll *UserLogic)SendCode(mobile string) error {
	err := ll.UserModel.SendCode(mobile)
	return err
}


func (ll *UserLogic)GetCode(mobile string) (typeuser.MobileCode, error) {
	code, err := ll.UserModel.GetCode(mobile)
	return typeuser.MobileCode{*code}, err
}

func (ll *UserLogic)Register(mobile, version string) (interface{}, error) {
	user, err := ll.UserModel.Register(mobile)

	accessToken, err := GenTokenTest("123456", map[string]interface{}{typeuser.JwtUserField: user.Mobile, typeuser.JwtVersionField: "v1.0.1"}, 10)
	if err == nil {
		user.AccessToken = accessToken
	}
	return user, err
}