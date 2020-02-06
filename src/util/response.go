package util

import (
	"demo_1/src/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinS struct {
	Ctx *gin.Context
}

type ResS struct {
	Code    constant.ResultCode `json:"code"`
	Message string              `json:"msg"`
	Data    interface{}         `json:"data"`
}

func (g *GinS) Response(code constant.ResultCode, msg string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, ResS{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
