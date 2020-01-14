package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"project_golang/common/baseresponse"
	uuid2 "project_golang/common/uuid"
	"project_golang/services/user/typeuser"
	"strconv"
	"time"
)

type UserModel struct {
	Biz        *redis.Client
	Collection *mongo.Collection
}

const MoBileCode = "MobileCode#%v" // 手机验证码可以
const UserField = "UserFields#%v"
const UserSaveKey = "UserSearchKeys"

func (mm *UserModel) FindUser(mobile string) (*typeuser.User, error) {
	val, err := mm.Biz.HGet(UserSaveKey, fmt.Sprintf(UserField, mobile)).Result()
	if err != nil {
		return nil, err
	}
	var user typeuser.User
	json.Unmarshal([]byte(val), &user)
	return &user, nil
}

func (mm *UserModel) FindAllUser() (*[]typeuser.User, error) {
	val, err := mm.Biz.HGetAll(UserSaveKey).Result()
	if err != nil {
		return nil, err
	}
	var users []typeuser.User
	for u := range val {
		userJson := val[u]
		var user typeuser.User
		json.Unmarshal([]byte(userJson), &user)
		users = append(users, user)
	}
	return &users, nil
}

func (mm *UserModel) Register(nickname, email, address, avatar, mobile string, age int64) (*typeuser.User, error) {
	user := typeuser.User{
		nil,
		nickname,
		mobile,
		email,
		age,
		address,
		avatar,
		nil,
	}
	uuid := uuid2.CreateUUID()
	user.ID = &uuid
	if user.ID == nil {
		return nil, &baseresponse.LysError{"创建用户 ID 失败"}
	}

	return mm.AddUser(nickname, email, address, avatar, mobile, *user.ID, age)
}

func (mm *UserModel) AddUser(nickname, email, address, avatar, mobile, uuid string, age int64) (*typeuser.User, error) {
	user := typeuser.User{
		&uuid,
		nickname,
		mobile,
		email,
		age,
		address,
		avatar,
		nil,
	}
	bty, _ := json.Marshal(user)
	cmd := mm.Biz.HSet(UserSaveKey, fmt.Sprintf(UserField, user.Mobile), string(bty))
	_, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*删除用户信息*/
func (mm *UserModel) DeleteUser(mobile string) error {
	_, err := mm.Biz.HDel(UserSaveKey, fmt.Sprintf(UserField, mobile)).Result()
	return err
}

/*编辑用户信息*/
func (mm *UserModel) EditUser(nickname, email, address, avatar, mobile, token string, age int64) (*typeuser.User, error) {

	usr, err := mm.FindUser(mobile)
	if err != nil {
		return nil, &baseresponse.LysError{"该用户不存在"}
	}
	usr, err = mm.AddUser(nickname, email, address, avatar, mobile, token, age)
	if err != nil {
		return nil, err
	}
	return usr, err
}

func (mm *UserModel) SendCode(mobile string) error {
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

func (mm *UserModel) GetCode(mobile string) (*string, error) {
	val, err := mm.Biz.Get(fmt.Sprintf(MoBileCode, mobile)).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}
