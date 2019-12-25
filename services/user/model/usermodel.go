package model

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"project_golang/services/user/typeuser"
	"time"
)

type UserModel struct {
	Biz *redis.Client
}

func (mm *UserModel)FindUser(mobile string) (*typeuser.User, error) {
	//user := typeuser.User{
	//	"alexluan",
	//	"18817322111",
	//	"2222@qq.com",
	//	12,
	//	"shanghai",
	//	"http://xxx.jpg",
	//}
	val, err := mm.Biz.Get(mobile).Result()
	if err != nil {
		return nil, err
	}
	var user typeuser.User
	json.Unmarshal([]byte(val), &user)
	return &user, nil
}

func (mm *UserModel)Register(mobile string) (*typeuser.User, error) {
	user := typeuser.User{
		"alexluan",
		mobile,
		"2222@qq.com",
		12,
		"shanghai",
		"http://xxx.jpg",
		nil,
	}

	bty, _ := json.Marshal(user)

	err := mm.Biz.Set(mobile, string(bty), 10*time.Second).Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}