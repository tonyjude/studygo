package process2

import (
	"encoding/json"
	"fmt"
	"net"

	"../../message"
	"../model"
	"../utils"
)

//UserProcess
type UserProcess struct {
	Conn   net.Conn
	UserId int
}

//通知所有在线用户
func (_this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}

		up.NotifyMeOnline(userId)
	}
}

func (_this *UserProcess) NotifyMeOnline(userId int) {
	//组织消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json Marshal err = ", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: _this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err = ", err)
		return
	}
}

func (_this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误！"
		}
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	resMes.Data = string(data)

	//对resMes 进行序列化， 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: _this.Conn,
	}

	err = tf.WritePkg(data)
	return
}

//ServerProcessLogin
func (_this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	// if loginMes.UserID == 100 && loginMes.UserPWD == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在"
	// }

	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPWD)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 500
			loginResMes.Error = "服务器内部错误。。。"
		}
	} else {
		loginResMes.Code = 200
		_this.UserId = loginMes.UserID
		userMgr.AddOnlineUser(_this)
		_this.NotifyOthersOnlineUser(loginMes.UserID)
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println(user, "登录成功!")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	resMes.Data = string(data)

	//对resMes 进行序列化， 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: _this.Conn,
	}

	err = tf.WritePkg(data)
	return
}
