package email

import (
	"demo_1/src/database"
	"demo_1/src/model"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	var email = model.Email{}
	if err := c.BindJSON(&email); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	if database.DB.Where(email.UserID).Find(&model.User{}).Error != nil {
		utilGin.Response(-1, "用户不存在", nil)
		return
	}
	var result = database.DB.Create(&email)
	if err := result.Error; err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	utilGin.Response(1, "添加成功", result.Value)
}

func Index(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	result := database.DB.Where(c.Param("id")).First(&model.Email{})
	if result.Error != nil {
		utilGin.Response(-1, result.Error.Error(), nil)
		return
	}
	utilGin.Response(1, "查询成功", result.Value)
}

func View(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	result := database.DB.Where("user_id = ?", c.Param("user")).Find(&[]model.Email{})
	if result.Error != nil {
		utilGin.Response(1, result.Error.Error(), nil)
		return
	}
	utilGin.Response(1, "查询成功", result.Value)
}

func Delete(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var email = model.Email{}
	result := database.DB.Where(c.Param("id")).Find(&email).Delete(c.Param("id"))
	if result.Error != nil {
		utilGin.Response(-1, result.Error.Error(), nil)
		return
	}
	utilGin.Response(1, "删除成功", result.Value)
}

func Update(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var email = model.Email{}
	if err := c.BindJSON(&email); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	result := database.DB.Where(email.ID).Find(&email).Update(&email)
	if result.Error != nil {
		utilGin.Response(-1, result.Error.Error(), nil)
		return
	}
	utilGin.Response(1, "更新成功", result.Value)
}
