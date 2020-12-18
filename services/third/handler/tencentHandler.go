package handler

import (
	"example.com/m/services/third/logic"
	"github.com/gin-gonic/gin"
)

type TencentHandler struct {
	TencentLogic logic.TencentLogic
}

func (handler TencentHandler) GetTencentInfo(ctx *gin.Context) {

}
