package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err", err)
		return nil
	}

	//返回对象
	client.conn = conn

	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1,PUBLIC MODEL")
	fmt.Println("2,PRIVATE MODEL")
	fmt.Println("3,UPDATE USERNAME")
	fmt.Println("0,EXIT")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("plase input legal scope number!")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("Plase input username:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return false
	}

	return true
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	fmt.Println("Plase input chat message , input 'exit' exit!")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (clinet *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := clinet.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println("Plase input chat user name , input 'exit' exit!")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("Plase input chat content , input 'exit' exit!")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			//发给服务器
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println("Plase input chat user name , input 'exit' exit!")
		fmt.Scanln(&remoteName)

	}
}

// 处理server回应的消息，直接显示标准输入输出即可
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn) //永久阻塞
}

func (clinet *Client) Run() {
	for clinet.flag != 0 {
		for clinet.menu() != true {
		}

		switch clinet.flag {
		case 1:
			clinet.PublicChat()
			break
		case 2:
			clinet.PrivateChat()
			break
		case 3:
			clinet.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server IP address (default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "set server port (default: 8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("client conn err!")
		return
	}

	//单独开启一个goroutine去处理Server的回执消息
	go client.DealResponse()

	fmt.Println("client conn scuuess!")
	client.Run()

	//select {}
}
