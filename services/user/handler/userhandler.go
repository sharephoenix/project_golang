package handler

import (
	"github.com/gin-gonic/gin"
	"project_golang/common/baserequest"
	"project_golang/common/baseresponse"
	"project_golang/services/user/logic"
	"project_golang/services/user/typeuser"
)

type UserHandler struct {
	Logic logic.UserLogic
}

type ReqUser struct {
	Mobile string `json:"mobile"` //`form:"mobile" json:"mobile" xml:"mobile" binding:"mobile"`
}

type LoginReq struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

/*获取用户信息*/
func (ll *UserHandler) GetUser(context *gin.Context) {
	mobile := context.Request.Header["Token"]
	if len(mobile) <= 0 {
		resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"lost token"})
		context.JSON(200, resp)
	}
	res, err := ll.Logic.GetUser(mobile[0])
	resp := baseresponse.ConvertGinResonse(res, err)
	context.JSON(200, resp)
}

/*注册用户信息*/
func (ll *UserHandler) Register(accessSecret string) func(*gin.Context) {
	return func(context *gin.Context) {
		version := context.Request.Header["Version"]

		var reqUser ReqUser
		baserequest.GetBody(context, &reqUser)

		accessToken, err := logic.GenTokenTest(accessSecret, map[string]interface{}{typeuser.JwtUserField: reqUser.Mobile, typeuser.JwtVersionField: "v1.0.1"}, 100000000000000000)
		if err == nil {
			context.Request.Header["Authorization"] = []string{accessToken}
		}

		if len(version) < 1 {
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"has no version in headers"})
			context.JSON(200, resp)
			return
		} else {
			res, err := ll.Logic.Register(reqUser.Mobile, version[0])
			res.AccessToken = accessToken
			resp := baseresponse.ConvertGinResonse(res, err)
			context.JSON(200, resp)
		}
	}
}

/*发送验证码*/
func (ll *UserHandler) SendCode(context *gin.Context) {
	context.Request.Header["Test-Header"] = []string{"TESSSSS"}
	mobile := context.Param("mobile")
	err := ll.Logic.SendCode(mobile)
	resp := baseresponse.ConvertGinResonse(nil, err)
	context.JSON(200, resp)
}

/*获取验证码*/
func (ll *UserHandler) GetCode(context *gin.Context) {
	mobile := context.Param("mobile")
	res, err := ll.Logic.GetCode(mobile)
	resp := baseresponse.ConvertGinResonse(res, err)
	context.JSON(200, resp)
}

/*登录*/
func (ll *UserHandler) Login(secretKey string) func(ctx *gin.Context) {
	return func(context *gin.Context) {
		var loginReq LoginReq
		baserequest.GetBody(context, &loginReq)
		if loginReq.Mobile == "" || loginReq.Code == "" {
			resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"参数错误"})
			context.JSON(200, resp)
			return
		}
		res, err := ll.Logic.Login(secretKey, loginReq.Mobile, loginReq.Code)
		println("000000")
		resp := baseresponse.ConvertGinResonse(res, err)
		context.JSON(200, resp)
	}
}

func (ll *UserHandler) Test(context *gin.Context) {
	resp := baseresponse.ConvertGinResonse(nil, &baseresponse.LysError{"test"})
	context.JSON(200, resp)
}
