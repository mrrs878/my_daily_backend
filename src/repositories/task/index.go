package task

import (
	"demo_1/src/constant"
	"demo_1/src/database"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
	"log"
)

func Add(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	var task = model.Task{}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	task.UserID = userId
	if err := c.BindJSON(&task); err != nil {
		log.Println(err.Error())
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if database.DB.Where(task.UserID).Find(&model.User{}).Error != nil {
		utilGin.Response(constant.FAILED, "用户不存在", nil)
		return
	}
	var result = database.DB.Create(&task)
	if err := result.Error; err != nil {
		log.Println(err)
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	utilGin.Response(constant.SUCCESS, "添加成功", result.Value)
}

func Index(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	result := database.DB.Where(c.Param("id")).First(&model.Task{})
	if result.Error != nil {
		utilGin.Response(constant.FAILED, result.Error.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "查询成功", result.Value)
}

func View(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	result := database.DB.Where("user_id = ?", userId).Find(&[]model.Task{})
	if result.Error != nil {
		utilGin.Response(constant.FAILED, result.Error.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "查询成功", result.Value)
}
