package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"project_golang/services/user/typeuser"
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
