package goods

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/goods"
	"demo_1/src/tool"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func GetGoodsByClassAndPage(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	page, err := tool.String2Uint(c.Param("page"))
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	class, err := tool.String2Uint(c.Param("class"))
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	var result []model.Goods
	if err := goods.ViewByClassAndPage(page, class, &result); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "获取成功", result)
}

func CreateGoods(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	addForm := types.CreateGoodsForm{}
	if err := c.BindJSON(&addForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_goods := model.Goods{
		BaseModel: model.BaseModel{
			CreateId: userId,
		},
		Name:        addForm.Name,
		Class:       addForm.Class,
		Service:     addForm.Service,
		Description: addForm.Description,
	}
	if err := goods.Add(&_goods); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "添加成功", _goods)
}

func UpdateGoods(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	updateForm := types.UpdateGoodForm{}
	if err := c.BindJSON(&updateForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_goods := model.Goods{
		BaseModel: model.BaseModel{
			UpdateId: userId,
		},
		Name:        updateForm.Name,
		Class:       updateForm.Class,
		Service:     updateForm.Service,
		Description: updateForm.Description,
	}
	if err := goods.Update(&_goods); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", _goods)
}

func DeleteGoods(c *gin.Context) {}
