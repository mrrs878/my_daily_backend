package msg

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/message"
	"demo_1/src/repositories/user"
	"demo_1/src/service/ws"
	"demo_1/src/util"
	"demo_1/src/util/pushSub"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type wsMsg struct {
	Label string `json:"label"`
	Data  string `json:"data"`
}

func pushMsgList(params interface{}) {
	client := params.(*ws.Client)

	msg := wsMsg{}
	var msgList []model.Message
	userId, err := strconv.ParseUint(client.Id, 10, 64)
	if err != nil {
		ws.WebSocketManager.SendAll([]byte(err.Error()))
		return
	}
	err = viewMsgByUser(uint(userId), &msgList)
	if err != nil {
		ws.WebSocketManager.SendAll([]byte(err.Error()))
		return
	}
	msg.Label = "pushMsgList"
	data, err := json.Marshal(msgList)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		return
	}
	msg.Data = string(data)
	res, err := json.Marshal(&msg)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		return
	}

	err = user.Update(&model.User{WsGroup: client.Group, WsId: client.Id}, "name = ?", client.Id)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		return
	}
	ws.WebSocketManager.Send(client.Id, client.Group, res)
}
func pushMsg(param *model.Message) {
	msg := wsMsg{}
	msg.Label = "pushMsg"
	data, err := json.Marshal(param)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		return
	}
	msg.Data = string(data)
	res, err := json.Marshal(&msg)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		return
	}
	ws.WebSocketManager.Send(strconv.FormatInt(int64(param.UserId), 10), "msg", res)
}

func Setup() {
	suber := pushSub.Suber{
		Label: "clientConnected",
		Cb:    pushMsgList,
	}
	suber.Sub("clientConnected")
}

func CreateMessage(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}

	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_msg := model.Message{}
	if err := c.ShouldBindJSON(&_msg); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	_msg.CreateId = userId
	if err := message.Add(&_msg); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "创建成功", _msg)
	pushMsg(&_msg)
}

func UpdateMsg(c *gin.Context) {}

func DeleteMsg(c *gin.Context) {}

func viewMsgByUser(userId uint, msgList *[]model.Message) error {
	err := message.ViewWithCondition(msgList, "user_id = ?", userId)
	return err
}

func ViewMsgByUser(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	var _msgList []model.Message
	err = viewMsgByUser(userId, &_msgList)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}

	utilGin.Response(constant.SUCCESS, "查询成功", _msgList)
}
