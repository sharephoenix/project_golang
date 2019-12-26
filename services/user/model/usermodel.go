package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"project_golang/common/baseresponse"
	"project_golang/services/user/typeuser"
	"strconv"
	"time"
)

type UserModel struct {
	Biz *redis.Client
}

const MoBileCode = "MobileCode#%v" // 手机验证码可以

func (mm *UserModel)FindUser(mobile string) (*typeuser.User, error) {
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

	err := mm.Biz.Set(mobile, string(bty), 3600 * 24 *time.Second).Err()
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (mm *UserModel)SendCode(mobile string) error {
	rand.Seed(time.Now().UnixNano())
	var code string
	for i := 0; i < 4; i++ {
		randNum := rand.Intn(9)
		code += strconv.Itoa(randNum)
	}
	if code == "" {
		return &baseresponse.LysError{"生成验证码失败"}
	}
	err := mm.Biz.Set(fmt.Sprintf(MoBileCode, mobile), code, 60*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}


func (mm *UserModel)GetCode(mobile string) (*string, error) {
	val, err := mm.Biz.Get(fmt.Sprintf(MoBileCode, mobile)).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}