package controller

import (
	"demo_1/src/middleware"
	"demo_1/src/repositories/email"
	"demo_1/src/service/auth"
	"demo_1/src/service/user"
	"demo_1/src/tool"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.SetupLogger(), middleware.SetUpException())

	engine.Static("assets", "./assets")
	engine.NoRoute(func(c *gin.Context) {
		utilGin := tool.GinS{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(context *gin.Context) {
		utilGin := tool.GinS{Ctx: context}
		utilGin.Response(1, "pong", nil)
	})

	EmailRouter := engine.Group("/email")
	EmailRouter.Use(middleware.JWTAuth())
	{
		EmailRouter.POST("/", email.Add)
		EmailRouter.PUT("/", email.Update)
		EmailRouter.GET("/", email.View)
		EmailRouter.GET("/:id", email.Index)
		EmailRouter.DELETE("/:id", email.Delete)
	}

	UserRouter := engine.Group("/user")
	UserRouter.Use(middleware.JWTAuth())
	{
		UserRouter.GET("/", user.GetInfo)
		UserRouter.DELETE("/:id", user.WriteOff)
	}

	AuthRouter := engine.Group("/auth")
	{
		AuthRouter.POST("/login", auth.Login)
		AuthRouter.POST("/register", auth.Register)
	}
}
