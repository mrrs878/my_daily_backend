package sw

import (
	"demo_1/src/constant"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
	"log"
)

func Subscribable(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	log.Println(c.Request.Body)

	utilGin.Response(constant.SUCCESS, "订阅成功", 1)
}

func PushMessage(c *gin.Context) {}
