package baseresponse

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct{
	Code int64 `json:"code"`
	Msg string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type LysError struct {
	Msg string `json:"msg"`
}

func (err *LysError)Error() string {
	return err.Msg
}
func FormatResponse(data interface{}, err error) string {
	resp := ErrorResponse {
		0,
		err.Error(),
		data,
	}
	if err != nil {
		resp.Code = -1
		bty, _ := json.Marshal(resp)
		return string(bty)
	}
	byt, err := json.Marshal(resp)
	if err == nil {
		return string(byt)
	}
	resp.Code = -1
	resp.Msg = err.Error()
	bty, _ := json.Marshal(resp)
	return string(bty)
}

func ConvertGinResonse(data interface{}, err error) map[string]interface{} {
	resp := ErrorResponse{
		0,
		"",
		data,
	}
	if err != nil {
		resp.Code = -1
		resp.Data = nil
		resp.Msg = err.Error()
	}
	if data == nil {
		resp.Code = -2
		resp.Msg = "返回值为空"
	}
	bty, _ := json.Marshal(resp)
	var m gin.H
	json.Unmarshal(bty, &m)
	return m
}
