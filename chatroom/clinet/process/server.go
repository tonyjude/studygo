package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"../../message"
	"../utils"
)

func ShowMenu(userId int) {
	fmt.Printf("-------------------恭喜%d登录成功-------------------\n", userId)
	fmt.Println("------------------- 1 显示在线用户列表-------------------")
	fmt.Println("------------------- 2 发送消息-------------------")
	fmt.Println("------------------- 3 信息列表-------------------")
	fmt.Println("------------------- 4 退出系统-------------------")

	var key int
	var content string
	sms := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入消息内容：")
		fmt.Scanf("%s\n", &content)
		sms.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统成功！")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确！")
	}
}

//和服务器端保持连接
func ServerProcessMes(conn net.Conn) {
	//不停读取服务器的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Println("client reading from server ....")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		//fmt.Printf("message = %v\n", mes)

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//处理服务端消息
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsMesType:
			var smsMes message.SmsMes
			json.Unmarshal([]byte(mes.Data), &smsMes)
			getGroupMes(&smsMes)
		default:
			fmt.Println("服务端未知消息！")
		}

	}
}
