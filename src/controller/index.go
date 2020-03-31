package controller

import (
	"demo_1/src/constant"
	"demo_1/src/middleware"
	"demo_1/src/service/auth"
	"demo_1/src/service/dataDict"
	"demo_1/src/service/habit"
	"demo_1/src/service/msg"
	"demo_1/src/service/sw"
	"demo_1/src/service/task"
	"demo_1/src/service/user"
	"demo_1/src/service/ws"
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

	UserRouter := engine.Group("/user")
	UserRouter.Use(middleware.JWTAuth())
	{
		UserRouter.GET("", user.GetInfo)
		UserRouter.DELETE("", user.WriteOffSelf)
		UserRouter.DELETE("/:id", user.WriteOff)
	}

	TaskRouter := engine.Group("/task")
	TaskRouter.Use(middleware.JWTAuth())
	{
		TaskRouter.POST("", task.CreateTask)
		TaskRouter.PUT("", task.Update)
		TaskRouter.GET("", task.ViewByUser)
		TaskRouter.GET("/:id", task.GetInfo)
		TaskRouter.DELETE("/:id", task.Delete)
	}

	AuthRouter := engine.Group("/auth")
	{
		AuthRouter.POST("/login", auth.Login)
		AuthRouter.POST("/register", auth.Register)
		AuthRouter.GET("/github/:code", auth.LoginByGitHub)
	}

	SubscriptionRouter := engine.Group("/sw")
	SubscriptionRouter.Use(middleware.JWTAuth())
	{
		SubscriptionRouter.POST("/sub", sw.Subscribable)
		SubscriptionRouter.POST("/push", sw.PushMessage)
	}

	HabitRouter := engine.Group("/habit")
	HabitRouter.Use(middleware.JWTAuth())
	{
		HabitRouter.POST("", habit.CreateHabit)
		HabitRouter.PUT("", habit.UpdateHabit)
		HabitRouter.GET("/:id", habit.IndexHabit)
		HabitRouter.GET("", habit.ViewHabitsByUser)
		HabitRouter.DELETE("/:id", habit.DeleteHabit)
	}

	MsgRouter := engine.Group("/msg")
	{
		MsgRouter.GET("/:id", ws.WebSocketManager.WsClient)
		MsgRouter.GET("", msg.ViewMsgByUser)
		MsgRouter.POST("", msg.CreateMessage)
		MsgRouter.DELETE("/:id", msg.DeleteMsg)
		MsgRouter.PUT("/:id", msg.UpdateMsg)
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
