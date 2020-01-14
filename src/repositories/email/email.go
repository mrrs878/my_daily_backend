package email

import (
	"demo_1/src/database"
	"demo_1/src/models"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	//validate := validator.New()
	//_ = validate.RegisterValidation("NameValid", controller.NameValid)
	//
	//if err := validate.Struct(s); err != nil {
	//	utilGin.Response(-1, err.Error(), nil)
	//	return
	//}

	var email = models.Email{}
	err := c.BindJSON(&email)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
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
	var email = models.Email{}
	utilGin.Response(1, "查询成功", database.DB.Where("id = ?", c.Param("id")).First(&email).Value)
}
