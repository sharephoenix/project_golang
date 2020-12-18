package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//https://graph.qq.com/user/get_user_info?access_token=YOUR_ACCESS_TOKEN&oauth_consumer_key=YOUR_APP_ID&openid=YOUR_OPENID
type TencentLogic struct {
	DBName string
}

type TencentUserInfoReq struct {
	AccessToken string `json:"access_token"`
	AppId       string `json:"app_id"`
	Openid      string `json:"openid"`
}

type TencentUserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	Nickname        string `json:"nickname"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	Gender          string `json:"gender"`
	IsYellowVip     string `json:"is_yellow_vip"`
	level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

func (m TencentLogic) GetTencentUserInfo(req TencentUserInfoReq) TencentUserInfo {
	getUrl := getTencentUrl(req.AccessToken, req.AppId, req.Openid)
	result := Get(getUrl)
	var userInfo TencentUserInfo
	json.Unmarshal([]byte(result), &userInfo)
	return userInfo
}

// 腾讯三方登录，获取用户信息url
func getTencentUrl(accessToken, appId, openId string) string {
	baseUrl := "https://graph.qq.com/user/get_user_info"
	path := fmt.Sprintf("?access_token=%s&oauth_consumer_key=%s&openid=%s", accessToken, appId, openId)
	return baseUrl + path
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
