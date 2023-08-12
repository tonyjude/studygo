package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听当前User channel消息的 goroutine
	go user.ListenMessage()

	return user
}

func (this *User) Online() {

	//用户上线，将用户加入onlinemap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "user online")
}

func (this *User) Offline() {

	//用户下线，将用户从onlinemap移除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "user offline")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

func (this *User) DoMessage(msg string) {
	if msg == "who" {
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":online...\n"
			this.SendMsg(onlineMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式： rename|张三
		newName := strings.Split(msg, "|")[1]

		//判断newName 是否存在
		this.server.mapLock.Lock()

		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("用户名已存在")
		} else {
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this

			this.Name = newName
			this.SendMsg("用户名修改成功： " + newName)
		}

		this.server.mapLock.Unlock()

	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式： to|张三|你好啊

		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用\"to|张三|你好啊\" \n")
			return
		}

		//获取对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在\n")
			return
		}

		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重新发送\n")
			return
		}

		remoteUser.SendMsg(this.Name + "对您说：" + content)
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User的channel方法，一旦有消息就返回给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
