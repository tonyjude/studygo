package main

import (
	"fmt"
	"os"

	"../process"
)

var userID int
var userPWD string
var userName string

func main() {
	var key int
	for true {
		fmt.Println("-------------------欢迎登陆多人聊天系统-------------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPWD)

			up := &process.UserProcess{}
			up.Login(userID, userPWD)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPWD)
			fmt.Println("请输入用户昵称")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userID, userPWD, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}

	}
}
