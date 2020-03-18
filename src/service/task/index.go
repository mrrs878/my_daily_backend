package task

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/task"
	"demo_1/src/tool"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	createForm := types.CreateTaskForm{}
	if err := c.ShouldBindJSON(&createForm); err != nil {
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
		Status: createForm.Status,
		UserId: userId,
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

func GetInfo(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_task := model.Task{}
	if err := task.Index(&_task); err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
	}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "获取成功", types.TaskInfoForm{
		Model:     _task.Model,
		Title:     _task.Title,
		Label:     _task.Label,
		Detail:    _task.Detail,
		AlarmTime: _task.AlarmTime,
		Status:    _task.Status,
		UserId:    userId,
	})
}

func ViewByUser(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	var _tasks []model.Task
	err = task.ViewWithCondition(&_tasks, "user_id = ?", userId)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "查询成功", _tasks)
}

func Delete(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	_task := model.Task{}
	var err error = nil
	if _task.ID, err = tool.String2Uint(c.Param("id")); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if _task.DeleteId, err = functions.GetUserIDFormContext(c); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if err = task.Del(&_task); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "删除成功", nil)
}

func Update(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var _task = model.Task{}
	var err error = nil
	if err = c.BindJSON(&_task); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if _task.UpdateId, err = functions.GetUserIDFormContext(c); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if err := task.Update(&_task); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", _task)
}
