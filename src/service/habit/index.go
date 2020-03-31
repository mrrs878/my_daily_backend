package habit

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/habit"
	"demo_1/src/tool"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func CreateHabit(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	createForm := types.CreateHabitForm{}
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_habit := model.Habit{
		Title:      createForm.Title,
		Label:      createForm.Label,
		Detail:     createForm.Detail,
		AlarmTime:  createForm.AlarmTime,
		AlarmDate:  createForm.AlarmDate,
		SuccessCnt: 0,
		FailedCnt:  0,
		UserId:     userId,
		BaseModel: model.BaseModel{
			CreateId: userId,
		},
	}
	if err := habit.Add(&_habit); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "创建成功", _habit)
}

func UpdateHabit(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var _habit = model.Habit{}
	var err error = nil
	if err = c.BindJSON(&_habit); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if _habit.UpdateId, err = functions.GetUserIDFormContext(c); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if err := habit.Update(&_habit); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", _habit)
}

func DeleteHabit(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	_habit := model.Habit{}
	var err error = nil
	if _habit.ID, err = tool.String2Uint(c.Param("id")); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if _habit.DeleteId, err = functions.GetUserIDFormContext(c); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	if err = habit.Del(&_habit); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "删除成功", nil)
}

func ViewHabitsByUser(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	var _habits []model.Habit
	err = habit.ViewWithCondition(&_habits, "user_id = ?", userId)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "查询成功", _habits)
}

func IndexHabit(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_habit := model.Habit{}
	if err := habit.Index(&_habit); err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
	}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "获取成功", types.HabitInfoForm{
		Model:      _habit.Model,
		Title:      _habit.Title,
		Label:      _habit.Label,
		Detail:     _habit.Detail,
		AlarmTime:  _habit.AlarmTime,
		AlarmDate:  _habit.AlarmDate,
		SuccessCnt: _habit.SuccessCnt,
		FailedCnt:  _habit.FailedCnt,
		UserId:     userId,
	})
}
