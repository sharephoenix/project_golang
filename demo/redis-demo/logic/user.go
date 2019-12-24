package logic

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
)

var (
	userToken =  "%s:%s"
)

type Error struct {
	Content string
}


type User struct {
	Account string 		`json:"account"`
	Password string 	`json:"password"`
	Name string			`json:"name"`
	Address string		`json:"address"`
	Sex int64			`json:"sex"`
}

/// 用户注册， 模块
func Register(account string, password string, name string, address string, sex int64) *Error {
	user := User{
		account,
		password,
		name,
		address,
		sex,
	}
	client := getRedis()
	jsonBytes, jsonerror := json.Marshal(user)
	if jsonerror != nil {
		return &Error{jsonerror.Error()}
	}
	err := client.Set(account, string(jsonBytes), 0).Err()
	if err == nil {
		fmt.Printf("插入数据成功")
		return nil
	}
	return &Error{err.Error()}
}

/// 检查用户是否注册
func CheckRegister(account string) (string, bool) {
	client := getRedis()
	defer client.Close()

	msg, err := client.Get(account).Result()
	fmt.Printf("result:%v - %v\n", msg, account)
	if err == nil {
		return msg, false
	}
	return msg, true
}

func Login(account string, password string) *User {
	client := getRedis()
	defer client.Close()

	val, err := client.Get(account).Result()
	if err == nil {
		var userr User
		e := json.Unmarshal([]byte(val), &userr)
		if e == nil && userr.Password == password {
			return &userr
		}
	}
	return nil
}

func CreateToken() (*string, *string) {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, nil
	}

	token := GetMd5String(base64.URLEncoding.EncodeToString(u2.Bytes()))
	return &token, nil
}

func getRedis() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return *client
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}