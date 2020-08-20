package logic

import (
	"encoding/json"
	"example.com/m/common/baseresponse"
	"example.com/m/demo/email"
	"example.com/m/services/user/model"
	"example.com/m/services/user/typeuser"
	"fmt"
	"log"
)

type UserLogic struct {
	UserModel model.UserModel
}

func (ll *UserLogic) GetUser(mobile string) (typeuser.User, error) {
	user, err := ll.UserModel.MgoFindUser(mobile)
	return *user, err
}

func (ll *UserLogic) SendCode(mobile string) error {
	code, err := ll.UserModel.CreateLoginCode(mobile)
	user, err := ll.GetUser(mobile)
	if err != nil {
		return err
	}
	fmt.Println(code, user.Email)
	ll.sendEmail(user.Email, *code)
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
	cacheUser, err := ll.UserModel.MgoFindUser(mobile)
	if err == nil {
		return nil, &baseresponse.LysError{"该用户已经存在" + cacheUser.Nickname}
	}
	user, err := ll.UserModel.MgoRegister(nickname, email, address, avatar, mobile, age)
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
	user, err := ll.UserModel.MgoFindUser(mobile)
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
	users, err := ll.UserModel.MgoFindAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ll *UserLogic) DeleteUser(mobile string) error {
	err := ll.UserModel.MgoDeleteUser(mobile)
	return err
}

func (ll *UserLogic) EditUser(nickname, email, address, avatar, mobile, token string, age int64) (*typeuser.User, error) {
	user, err := ll.UserModel.MgoEditUser(nickname, email, address, avatar, mobile, token, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

/// 初始化 admin 用户
func (ll *UserLogic) InitilizeAdmin() {
	user, err := ll.Register("alex", "326083325@qq.com", "shanghai", "", "18817322818", 21, "1.1.1")
	if err != nil {
		fmt.Println("初始化用户失败", err)
	} else {
		bty, err0 := json.Marshal(user)
		if err0 == nil {
			fmt.Println("初始用户为：", string(bty))
		} else {
			fmt.Println("初始用户为: 解析错误")
		}
	}
	user0, er := ll.UserModel.MgoFindUser("18817322818")
	if er == nil {
		bty, err0 := json.Marshal(user0)
		if err0 == nil {
			fmt.Println("获取用户：", string(bty))
		} else {
			fmt.Println("获取用户: 解析错误")
		}
	} else {
		fmt.Println("获取用户失败", er)
	}
}

func (ll *UserLogic) sendEmail(to, code string) {
	//定义收件人
	mailTo := []string{
		to,
	}
	//邮件主题为"Hello"
	subject := "验证码"
	// 邮件正文
	body := "验证码为：" + code

	err := email.SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("send fail")
		return
	}
	fmt.Println("send successfully")
}
