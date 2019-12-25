package handler

import (
	"github.com/gin-gonic/gin"
	"project_golang/common/baserequest"
	"project_golang/common/baseresponse"
	"project_golang/services/user/logic"
)

type UserHandler struct {
	Logic logic.UserLogic
}

type ReqUser struct {
	Mobile string `json:"mobile"`//`form:"mobile" json:"mobile" xml:"mobile" binding:"mobile"`
}

func (ll *UserHandler)GetUser(context *gin.Context) {
	mobile := context.Param("mobile")
	res, err := ll.Logic.GetUser(mobile)
	resp := baseresponse.ConvertGinResonse(res, err)
	context.JSON(200, resp)
}

func (ll *UserHandler)Register(context *gin.Context) {
	var reqUser ReqUser
	baserequest.GetBody(context, &reqUser)
	res, err := ll.Logic.Register(reqUser.Mobile)
	resp := baseresponse.ConvertGinResonse(res, err)
	context.JSON(200, resp)
}