package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinS struct {
	Ctx *gin.Context
}

type ResS struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *GinS) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, ResS{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
