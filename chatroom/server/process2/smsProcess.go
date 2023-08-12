package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"../../message"
	"../utils"
)

//SmsProcess
type SmsProcess struct {
}

//向所有在线用户发送消息
func (_this *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}

		_this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (_this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser err = ", err)
		return
	}
}
