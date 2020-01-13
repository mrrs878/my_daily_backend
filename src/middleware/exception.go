package middleware

import (
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
)

func SetUpException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				utilGin := util.GinS{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
