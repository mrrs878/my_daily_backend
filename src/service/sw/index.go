package sw

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/subscription"
	"demo_1/src/types"
	"demo_1/src/util"
	"github.com/gin-gonic/gin"
)

func Subscribable(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_subForm := types.SubscriptionForm{}
	if err := c.ShouldBindJSON(&_subForm); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_sub := model.Subscription{}
	_sub.UserId = userId
	_sub.Endpoint = _subForm.Endpoint
	_sub.ApplicationServerKey = _subForm.ApplicationServerKey
	_sub.Auth = _subForm.Auth
	_sub.ExpirationTime = _subForm.ExpirationTime
	_res, err := subscription.Add(&_sub)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "订阅成功", _res)
}

func PushMessage(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	_msgForm := types.PushMsgForm{}
	err := c.ShouldBind(&_msgForm)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	//_msg := model.PushMsg{}
	//_msg.UserId = _msgForm.UserId
	//_msg.Detail = _msgForm.Detail
	//log.Println(_msg.Detail)
	var _subs []model.Subscription
	err = subscription.ViewWithCondition(&_subs, "user_id = ?", _msgForm.UserId)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "推送成功", _subs)
}
