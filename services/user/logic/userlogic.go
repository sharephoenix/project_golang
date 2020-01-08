package logic

import (
	"project_golang/common/baseresponse"
	"project_golang/services/user/model"
	"project_golang/services/user/typeuser"
)

type UserLogic struct {
	UserModel model.UserModel
}

func (ll *UserLogic) GetUser(mobile string) (interface{}, error) {
	user, err := ll.UserModel.FindUser(mobile)
	return user, err
}

func (ll *UserLogic) SendCode(mobile string) error {
	err := ll.UserModel.SendCode(mobile)
	return err
}

func (ll *UserLogic) GetCode(mobile string) (*typeuser.MobileCode, error) {
	code, err := ll.UserModel.GetCode(mobile)
	if err != nil {
		return nil, err
	}
	return &typeuser.MobileCode{*code}, err
}

func (ll *UserLogic) Register(nickname, email, address, avatar, mobile string, age int64, version string) (*typeuser.User, error) {
	cacheUser, err := ll.UserModel.FindUser(mobile)
	if err == nil {
		return nil, &baseresponse.LysError{"该用户已经存在" + cacheUser.Nickname}
	}
	user, err := ll.UserModel.Register(nickname, email, address, avatar, mobile, age)
	return user, err
}

func (ll *UserLogic) Login(secretKey, mobile, code string) (*typeuser.User, error) {
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
	token, err := GenTokenTest(secretKey, map[string]interface{}{"usr": mobile}, 3600*24)
	if err != nil {
		return nil, err
	}
	user.AccessToken = token
	return user, nil
}

func (ll *UserLogic) FindAll() (*[]typeuser.User, error) {
	users, err := ll.UserModel.FindAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ll *UserLogic) DeleteUser(mobile string) error {
	err := ll.UserModel.DeleteUser(mobile)
	return err
}

func (ll *UserLogic) EditUser(nickname, email, address, avatar, mobile, token string, age int64) (*typeuser.User, error) {
	user, err := ll.UserModel.EditUser(nickname, email, address, avatar, mobile, token, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
