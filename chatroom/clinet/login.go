package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../message"
)

func login(userID int, userPWD string) (err error) {
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
	loginMes.UserPWD = userPWD

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

	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("Login Success")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return

}
