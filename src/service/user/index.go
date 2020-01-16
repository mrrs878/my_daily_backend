package user

import (
	"demo_1/src/middleware"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	value, exists := c.Get("claims")
	if !exists {
		utilGin.Response(-1, "token失效", nil)
		return
	}
	result, ok := (value).(middleware.CustomClaims)
	if !ok {
		utilGin.Response(-1, "token失效", nil)
		return
	}
	_user := model.User{}
	_user.ID = result.ID
	if err := user.Index(&_user); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "获取成功", types.UserInfoForm{
		Model:  _user.Model,
		Name:   _user.Name,
		Emails: _user.Emails,
		Role:   _user.Role,
		Token:  c.Request.Header.Get("Authorization"),
	})
}
