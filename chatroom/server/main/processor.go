package main

import (
	"fmt"
	"io"
	"net"

	"../../message"
	"../process2"
	"../utils"
)

//处理器Processor
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数，根据客户端发送的消息种类不同，决定调用哪个函数
func (_this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: _this.Conn,
		}

		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: _this.Conn,
		}

		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		sp := &process2.SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不匹配！")
	}
	return
}

func (_this *Processor) process() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: _this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit, server exit;")
				return err
			}
			fmt.Println("readPkg err=", err)
			return err
		}

		err = _this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
