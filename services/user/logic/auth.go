package logic

import (
	"jwt-go"
	"project_golang/common/baseresponse"
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
	exp := time.Now().Add(time.Second * 3600 * 24).Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = now
	claims["nbf"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err := token.SignedString([]byte(secretKey))

	return jwtToken, err
}

// 解析里面的数据
func BackGenToken(jwtToken, secretKey string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
		return []byte(secretKey), nil
	})
	if err != nil{
		//fmt.Println("HS256的token解析错误，err:", err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &baseresponse.LysError{"ParseHStoken:claims类型转换失败"}
	}
	return claims, nil
}