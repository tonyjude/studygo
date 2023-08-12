package process

import (
	"fmt"

	"../../message"
	"../model"
)

var (
	onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
	CurUser     model.CurUser         //在用户登录成功后完成初始化
)

func updateUserStatus(notityUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notityUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notityUserStatusMes.UserId,
		}
	}

	user.UserStatus = notityUserStatusMes.Status
	onlineUsers[notityUserStatusMes.UserId] = user

	outputOnlineUser()
}

func outputOnlineUser() {
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		//如果不是自己
		fmt.Println("用户ID:\t", id)
	}
}
