package handler

import (
	"example.com/m/common/baserequest"
	"example.com/m/common/baseresponse"
	"example.com/m/services/user/logic"
	"example.com/m/services/user/typeuser"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	Logic logic.UserLogic
}

type ReqUser struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Avatar   string `json:"avatar"`
	Age      int64  `json:"age"`
	Mobile   string `json:"mobile"` //`form:"mobile" json:"mobile" xml:"mobile" binding:"mobile"`
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
			res, err := ll.Logic.Register(reqUser.Nickname, reqUser.Email, reqUser.Address, reqUser.Avatar, reqUser.Mobile, reqUser.Age, version[0])
			if err == nil {
				res.AccessToken = accessToken
			}
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

/*获取所有用户信息*/
func (ll *UserHandler) FindAll(context *gin.Context) {
	users, err := ll.Logic.FindAll()
	resp := baseresponse.ConvertGinResonse(users, err)
	context.JSON(200, resp)
}

/*编辑用户*/
func (ll *UserHandler) EditUser(context *gin.Context) {
	var reqUser ReqUser
	baserequest.GetBody(context, &reqUser)
	usr, err := ll.Logic.EditUser(reqUser.Nickname, reqUser.Email, reqUser.Address, reqUser.Avatar, reqUser.Mobile, reqUser.ID, reqUser.Age)
	resp := baseresponse.ConvertGinResonse(usr, err)
	context.JSON(200, resp)
}

/*删除用户*/
func (ll *UserHandler) DeleteUser(context *gin.Context) {
	type ReqDelete struct {
		Mobile string `json:"mobile"`
	}
	var reqDelete ReqDelete
	baserequest.GetBody(context, &reqDelete)
	err := ll.Logic.DeleteUser(reqDelete.Mobile)

	resp := baseresponse.ConvertGinResonse(nil, err)
	context.JSON(200, resp)
}

func (ll *UserHandler) Test(c *gin.Context) {
	response, err := http.Get("https://www.baidu.com")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//var a []interface{}
	//a = append(a, "alexluan")
	//bty, _ := json.Marshal(a)
	//context.String(200, "%v", bty)
}
