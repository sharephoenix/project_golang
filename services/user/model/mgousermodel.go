package model

import (
	"encoding/json"
	"example.com/m/common/baseresponse"
	"example.com/m/services/user/typeuser"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (mm *UserModel) MgoFindUser(mobile string) (*typeuser.User, error) {
	var user typeuser.User
	err := mm.Collection.FindOne(bson.M{"mobile": mobile}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (mm *UserModel) MgoFindAllUser() (*[]typeuser.User, error) {
	var users []typeuser.User
	err := mm.Collection.All(bson.M{}, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (mm *UserModel) MgoRegister(nickname, email, address, avatar, mobile string, age int64) (*typeuser.User, error) {
	data := bson.M{
		"nickname": nickname,
		"email":    email,
		"address":  address,
		"mobile":   mobile,
		"avatar":   avatar,
		"age":      age,
	}

	bty, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	mm.Collection.InsertOne(&data)
	var user typeuser.User
	json.Unmarshal(bty, &user)
	return &user, nil
}

/*编辑用户信息*/
func (mm *UserModel) MgoEditUser(nickname, email, address, avatar, mobile, token string, age int64) (*typeuser.User, error) {

	usr, err := mm.MgoFindUser(mobile)
	if err != nil {
		return nil, &baseresponse.LysError{"该用户不存在0"}
	}
	fmt.Println("eee", mobile, address, nickname, email, avatar, age)
	selector := bson.M{"mobile": mobile}
	data := bson.M{
		"nickname": nickname,
		"email":    email,
		"address":  address,
		"mobile":   mobile,
		"avatar":   avatar,
		"age":      age,
	}
	err = mm.Collection.Update(selector, data)
	if err != nil {
		fmt.Println("eee", err.Error())
		return nil, err
	}
	usr, err = mm.MgoFindUser(mobile)
	if err != nil {
		return nil, &baseresponse.LysError{"该用户不存在1"}
	}
	return usr, err
}

/*删除用户信息*/
func (mm *UserModel) MgoDeleteUser(mobile string) error {
	selector := bson.M{"mobile": mobile}
	err := mm.Collection.Delete(selector) //mm.Biz.HDel(UserSaveKey, fmt.Sprintf(UserField, mobile)).Result()
	return err
}
