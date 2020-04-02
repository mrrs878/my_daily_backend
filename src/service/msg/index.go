package msg

import (
	"demo_1/src/constant"
	"demo_1/src/functions"
	"demo_1/src/model"
	"demo_1/src/repositories/message"
	"demo_1/src/repositories/user"
	"demo_1/src/tool"
	"demo_1/src/types"
	"demo_1/src/util"
	"demo_1/src/util/pushSub"
	"demo_1/src/util/ws"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type wsMsg struct {
	Label string `json:"label"`
	Data  string `json:"data"`
}
type wsResMsg struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func pushMsgList(params interface{}) {
	client := params.(*ws.Client)

	defer func() {
		if err := recover(); err != nil {
			ws.WebSocketManager.SendAll(err.([]byte))
		}
	}()
	msg := wsMsg{}
	var msgList []model.Message
	userId, err := strconv.ParseUint(client.Id, 10, 64)
	if err != nil {
		panic(err)
	}
	err = viewsMsg(&msgList, "user_id=?", uint(userId))
	if err != nil {
		panic(err)
	}
	msg.Label = "pushMsgList"
	data, err := json.Marshal(msgList)
	if err != nil {
		panic(err)
	}
	msg.Data = string(data)
	res, err := json.Marshal(&msg)
	if err != nil {
		panic(err)
	}

	err = user.Update(&model.User{WsGroup: client.Group, WsId: client.Id}, "id = ?", uint(userId))
	if err != nil {
		panic(err)
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

var ReceivedHandlers = map[string]func(*wsMsg) ([]byte, error){
	"ReadAll": func(msgData *wsMsg) ([]byte, error) {
		_msg := model.Message{}
		var _msgList []model.Message
		if err := json.Unmarshal([]byte(msgData.Data), &_msg); err != nil {
			return nil, err
		}
		newVal := types.UpdateNewVal{
			"status": _msg.Status,
		}
		if err := updateMsg(&newVal, &_msgList, "user_id=?", _msg.UserId); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		data, err := json.Marshal(_msgList)
		return data, err
	},
	"ReadOne": func(msgData *wsMsg) ([]byte, error) {
		_newVal := types.UpdateNewVal{"status": 1}
		_msg := model.Message{}
		if err := json.Unmarshal([]byte(msgData.Data), &_msg); err != nil {
			return nil, err
		}
		_msg.Status = 1
		if err := updateMsg(&_newVal, &_msg, "id=?", _msg.ID); err != nil {
			return nil, err
		}
		data, err := json.Marshal(_msg)
		return data, err
	},
	"GetAll": func(msgData *wsMsg) ([]byte, error) {
		_msg := model.Message{}
		if err := json.Unmarshal([]byte(msgData.Data), &_msg); err != nil {
			return nil, err
		}
		var _msgList []model.Message
		err := viewsMsg(&_msgList, "user_id=?", _msg.UserId)
		data, err := json.Marshal(_msgList)
		return data, err
	},
	"RemoveOne": func(msgData *wsMsg) ([]byte, error) {
		_msg := model.Message{}
		if err := json.Unmarshal([]byte(msgData.Data), &_msg); err != nil {
			return nil, err
		}
		err := deleteMsg(_msg.ID, _msg.UserId)
		data, err := json.Marshal(struct{ ID uint }{
			ID: _msg.ID,
		})
		return data, err
	},
}

func sendResponseMsg(resData *wsMsg, msgData *ws.MessageData) {
	res, err := json.Marshal(&resData)
	if err != nil {
		log.Println("send ws msg error: ", err.Error())
		res = []byte(err.Error())
	}
	ws.WebSocketManager.Send(msgData.Id, msgData.Group, res)
}
func receivedMsg(params interface{}) {
	msgData := params.(*ws.MessageData)
	_msgData := wsMsg{}
	resData := wsMsg{}
	resDataData := wsResMsg{}
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			sendResponseMsg(&resData, msgData)
		}
	}()

	if err := json.Unmarshal(msgData.Message, &_msgData); err != nil {
		resDataData.Success = false
		resDataData.Data = err.Error()
		panic("")
	}

	fn := ReceivedHandlers[_msgData.Label]
	resData.Label = _msgData.Label
	if fn == nil {
		resDataData.Success = false
		resDataData.Data = ""
		panic("")
	}
	data, err := fn(&_msgData)
	if err != nil {
		resDataData.Success = false
		resDataData.Data = err.Error()
		panic("")
	}
	resDataData.Data = string(data)
	resDataData.Success = true
	tmp, err := json.Marshal(resDataData)
	resData.Data = string(tmp)
	sendResponseMsg(&resData, msgData)
}

func Setup() {
	connSuber := pushSub.Suber{
		Label: "msgClientConnected",
		Cb:    pushMsgList,
	}
	connSuber.Sub("msgClientConnected")
	msgSuber := pushSub.Suber{
		Label: "receivedMsg",
		Cb:    receivedMsg,
	}
	msgSuber.Sub("msgClientReceived")
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

func updateMsg(newVal *types.UpdateNewVal, out interface{}, args ...interface{}) error {
	err := message.Update(newVal, out, args[0], args[1:]...)
	return err
}
func UpdateMsg(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	_newVal := types.UpdateNewVal{}
	_msg := model.Message{}
	if err := c.ShouldBindJSON(&_newVal); err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	err := updateMsg(&_newVal, &_msg)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}
	utilGin.Response(constant.SUCCESS, "更新成功", nil)
}

func deleteMsg(id uint, userId uint) error {
	_message := model.Message{}
	var err error = nil
	_message.ID = id
	_message.DeleteId = userId
	if err = message.Del(&_message); err != nil {
		return err
	}
	return nil
}
func DeleteMsg(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	var msgId, userId uint
	var err error
	defer func() {
		if err := recover(); err != nil {
			utilGin.Response(constant.FAILED, err.(string), nil)
		}
	}()
	msgId, err = tool.String2Uint(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}
	userId, err = functions.GetUserIDFormContext(c)
	if err != nil {
		panic(err.Error())
	}
	if err = deleteMsg(msgId, userId); err != nil {
		panic(err.Error())
	}
	utilGin.Response(constant.SUCCESS, "删除成功", nil)
}

func viewsMsg(msgList *[]model.Message, args ...interface{}) error {
	err := message.ViewWithCondition(msgList, args[0], args[1:]...)
	return err
}
func ViewsMsg(c *gin.Context) {
	utilGin := util.GinS{Ctx: c}
	userId, err := functions.GetUserIDFormContext(c)
	if err != nil {
		utilGin.Response(constant.FAILED, err.Error(), nil)
		return
	}

	var _msgList []model.Message
	err = viewsMsg(&_msgList, "user_id=?", userId)
	if err != nil {
		utilGin.Response(constant.SUCCESS, err.Error(), nil)
		return
	}

	utilGin.Response(constant.SUCCESS, "查询成功", _msgList)
}
