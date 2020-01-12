package controller

import (
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func SetupRouter (engine *gin.Engine) {
	//engine.Use(logger.Setup(), exception)

	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.GinS{ Ctx:c }
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(context *gin.Context) {
		utilGin := response.GinS{ Ctx: context }
		utilGin.Response(1, "pong", nil)
	})

	ProductRouter := engine.Group("")
	{
		ProductRouter.POST("/product", product)
	}
}