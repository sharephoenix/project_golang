package baserequest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetBody(context *gin.Context, body interface{}) {
	buf := make([]byte, 1024)
	context.Request.Body.Read(buf)
	var buuf []byte
	for i :=0; i<len(buf); i++ {
		if buf[i] != 0 {
			buuf = append(buuf, buf[i])
		}
	}
	json.Unmarshal(buuf, &body)
}