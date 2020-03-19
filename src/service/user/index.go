package user

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/user"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInfo(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_user := model.User{}
	if err := functions.GetUserInfoFromContext(c, &_user); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
	}
	utilGin.Response(constant.SUCCESS, "获取成功", types.UserInfoForm{
		Model: _user.Model,
		Name:  _user.Name,
		Role:  _user.Role,
		Token: c.Request.Header.Get("Authorization"),
	})
}

func WriteOff(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_user := model.User{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32<<(^uint(0)>>63))
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
	}
	_user.ID = uint(id)
	if err := user.Delete(&_user); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "注销成功", nil)
}

func WriteOffSelf(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	var (
		id  uint
		err error
	)
	if id, err = functions.GetUserIDFormContext(c); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
	}

	_user := model.User{}
	_user.ID = id
	if err := user.Delete(&_user); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "注销成功", nil)
}
