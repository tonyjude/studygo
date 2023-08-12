package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"../../message"
	"../utils"
)

type UserProcess struct {
}

func (_this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1，链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//2，准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//把data赋给mes.Data字段
	mes.Data = string(data)

	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功！")
	} else {
		fmt.Println(registerResMes.Error)
	}

	os.Exit(0)
	return
}

func (_this *UserProcess) Login(userID int, userPwd string) (err error) {
	//1，链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	defer conn.Close()

	//2，准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPWD = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//把data赋给mes.Data字段
	mes.Data = string(data)

	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail", err)
		return
	}

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//time.Sleep(5 * time.Second)
	//fmt.Println("sleep 5 second...")
	//fmt.Printf("client send message length=%d, content=%s", len(data), string(data))

	tf := &utils.Transfer{
		Conn: conn,
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userID
		CurUser.UserStatus = message.UserOnline
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			if v == userID {
				continue
			}
			fmt.Println("用户id:\t", v)
			//客户端onlineUser
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		//开启一个协程，监听服务器数据
		go ServerProcessMes(conn)
		//显示登录成功的菜单
		for {
			ShowMenu(loginMes.UserID)
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return

}
