package user

import (
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"demo_1/src/tool"
	"demo_1/src/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInfo(c *gin.Context) {
	utilGin := tool.GinS{Ctx: c}

	_user := model.User{}
	if err := functions.GetUserInfoFromContext(c, &_user); err != nil {
		utilGin.Response(-1, err.Error(), nil)
	}
	utilGin.Response(0, "获取成功", types.UserInfoForm{
		Model:  _user.Model,
		Name:   _user.Name,
		Emails: _user.Emails,
		Role:   _user.Role,
		Token:  c.Request.Header.Get("Authorization"),
	})
}

func WriteOff(c *gin.Context) {
	utilGin := tool.GinS{Ctx: c}

	_user := model.User{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32<<(^uint(0)>>63))
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
	}
	_user.ID = uint(id)
	if err := user.Delete(&_user); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "注销成功", nil)
}
