package controller

import (
	"demo_1/src/constant"
	"demo_1/src/middleware"
	"demo_1/src/repositories/task"
	"demo_1/src/service/auth"
	"demo_1/src/service/dataDict"
	"demo_1/src/service/goods"
	"demo_1/src/service/user"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	engine.Use(middleware.SetUpCors(), middleware.SetupLogger(), middleware.SetUpException())

	engine.Static("assets", "./assets")
	engine.NoRoute(func(c *gin.Context) {
		utilGin := util.GinS{Ctx: c}
		utilGin.Response(constant.FAILED, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(context *gin.Context) {
		utilGin := util.GinS{Ctx: context}
		utilGin.Response(constant.SUCCESS, "pong", nil)
	})

	TaskRouter := engine.Group("/task")
	TaskRouter.Use(middleware.JWTAuth())
	{
		TaskRouter.POST("", task.Add)
		TaskRouter.PUT("", task.Update)
		TaskRouter.GET("", task.View)
		TaskRouter.GET("/:id", task.Index)
		TaskRouter.DELETE("/:id", task.Delete)
	}

	UserRouter := engine.Group("/user")
	UserRouter.Use(middleware.JWTAuth())
	{
		UserRouter.GET("/", user.GetInfo)
		UserRouter.DELETE("/", user.WriteOffSelf)
		UserRouter.DELETE("/:id", user.WriteOff)
	}

	AuthRouter := engine.Group("/auth")
	{
		AuthRouter.POST("/login", auth.Login)
		AuthRouter.POST("/register", auth.Register)
	}

	GoodsRouter := engine.Group("/goods")
	GoodsRouter.Use(middleware.JWTAuth())
	{
		GoodsRouter.GET("/class/:class/page/:page", goods.GetGoodsByClassAndPage)
		GoodsRouter.POST("/", goods.CreateGoods)
		GoodsRouter.PUT("/", goods.UpdateGoods)
	}

	DataDictRouter := engine.Group("/dataDict")
	DataDictRouter.Use(middleware.JWTAuth())
	{
		DataDictRouter.POST("/", dataDict.AddDataDict)
		DataDictRouter.PUT("/", dataDict.UpdateDataDict)
		DataDictRouter.GET("/", dataDict.ViewAllDataDict)
		DataDictRouter.DELETE("/:id", dataDict.WriteOff)
		DataDictRouter.GET("/group/:group", dataDict.ViewByGroupName)
	}
}
