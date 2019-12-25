package logic

import (
	"fmt"
	"jwt-go"
	"time"
)

/// 获取 jwt token
func GenTokenTest(secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
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