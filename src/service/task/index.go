package task

import (
	"demo_1/src/constant"
	"demo_1/src/database"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/task"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	createForm := types.CreateTaskForm{}
	if err := c.BindJSON(&createForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_task := model.Task{
		Title:  createForm.Title,
		Label:  createForm.Label,
		Detail: createForm.Detail,
		BaseModel: model.BaseModel{
			CreateId: userId,
		},
	}
	if err := task.Add(&_task); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "创建成功", _task)
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

func Delete(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var task = model.Task{}
	result := database.DB.Where(c.Param("id")).Find(&task).Delete(c.Param("id"))
	if result.Error != nil {
		utilGin.Response(constant.FAILED, result.Error.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "删除成功", result.Value)
}

func Update(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var task = model.Task{}
	if err := c.BindJSON(&task); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	result := database.DB.Where(task.ID).Find(&task).Update(&task)
	if result.Error != nil {
		utilGin.Response(constant.FAILED, result.Error.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", result.Value)
}
