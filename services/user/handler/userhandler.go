package handler

import (
	"github.com/gin-gonic/gin"
	"project_golang/common/baseresponse"
	"project_golang/services/user/logic"
)

type UserHandler struct {
	Logic logic.UserLogic
}

func (ll *UserHandler)GetUser(context *gin.Context) {
	res, err := ll.Logic.GetUser("188121")
	resp := baseresponse.ConvertGinResonse(res, err)
	context.JSON(200, resp)
}