package logic

import (
	"fmt"
	"jwt-go"
	"time"
)

//iss: 签发者
//
//sub: 面向的用户
//
//aud: 接收方
//
//exp: 过期时间
//
//nbf: 生效时间
//
//iat: 签发时间
//
//jti: 唯一身份标识

/// 获取 jwt token
func GenTokenTest(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now * 60 * 2
	claims["iat"] = now
	claims["nbf"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err := token.SignedString([]byte(secretKey))

	BackGenToken(jwtToken, secretKey)

	return jwtToken, err
}

// 解析里面的数据
func BackGenToken(jwtToken, secretKey string) jwt.MapClaims{

	token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
		return []byte(secretKey), nil
	})
	if err != nil{
		fmt.Println("HS256的token解析错误，err:", err)
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil
	}
	return claims
}