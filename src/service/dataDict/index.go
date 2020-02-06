package dataDict

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/dataDict"
	"demo_1/src/tool"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddDataDict(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	addForm := types.CreateDataDictForm{}
	if err := c.BindJSON(&addForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_dataDict := model.DataDict{
		GroupName: addForm.GroupName,
		Label:     addForm.Label,
		Value:     addForm.Value,
		BaseModel: model.BaseModel{
			CreateId: userId,
		},
	}
	if err := dataDict.Add(&_dataDict); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "创建成功", _dataDict)
}

func UpdateDataDict(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	updateForm := types.UpdateDataDictForm{}
	if err := c.BindJSON(&updateForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_dataDict := model.DataDict{
		Model: gorm.Model{
			ID: updateForm.Id,
		},
		GroupName: updateForm.GroupName,
		Label:     updateForm.Label,
		Value:     updateForm.Value,
		BaseModel: model.BaseModel{
			UpdateId: userId,
		},
	}
	if err := dataDict.Update(&_dataDict); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", _dataDict)
}

func WriteOff(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	_dataDict := model.DataDict{}
	id, err := tool.String2Uint(c.Param("id"))
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_dataDict.ID = id
	if err := dataDict.Delete(&_dataDict); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "删除成功", _dataDict)
}

func ViewAllDataDict(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var _dataDict []model.DataDict
	if err := dataDict.View(&_dataDict); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "获取成功", _dataDict)
}

func ViewByGroupName(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var _dataDict []model.DataDict

	condition := model.DataDict{
		GroupName: c.Param("group"),
	}
	if err := dataDict.ViewWithCondition(&condition, &_dataDict); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "获取成功", _dataDict)
}
