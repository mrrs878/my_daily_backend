package controller

import (
	"demo_1/src/middleware"
	"demo_1/src/repositories/email"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.SetupLogger(), middleware.SetUpException())

	engine.Static("assets", "./assets")
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.GinS{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(context *gin.Context) {
		utilGin := util.GinS{Ctx: context}
		utilGin.Response(1, "pong", nil)
	})

	//ProductRouter := engine.Group("")
	//{
	//	ProductRouter.POST("/product", product)
	//}

	EmailRouter := engine.Group("/email")
	{
		EmailRouter.POST("/", email.Add)
		EmailRouter.GET("/:id", email.Index)
	}
}
