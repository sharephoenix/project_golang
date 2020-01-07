package logic

import (
	"project_golang/common/baseresponse"
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


func (ll *UserLogic)GetCode(mobile string) (*typeuser.MobileCode, error) {
	code, err := ll.UserModel.GetCode(mobile)
	if err != nil {
		return nil, err
	}
	return &typeuser.MobileCode{*code}, err
}

func (ll *UserLogic)Register(mobile, version string) (*typeuser.User, error) {
	user, err := ll.UserModel.Register(mobile)
	return user, err
}

func (ll *UserLogic)Login(secretKey, mobile, code string) (*typeuser.User, error) {
	realCode, err := ll.UserModel.GetCode(mobile)
	if err != nil {
		return nil, err
	}
	if *realCode != code {
		return nil, &baseresponse.LysError{"验证码错误"}
	}
	user, err := ll.UserModel.FindUser(mobile)
	if err != nil {
		return nil, &baseresponse.LysError{typeuser.NETERROR_NO_USER}
	}
	token, err := GenTokenTest(secretKey, map[string]interface{}{"usr": mobile}, 3600 * 24)
	if err != nil {
		return nil, err
	}
	user.AccessToken = token
	return user, nil
}