package main

import (
	"fmt"
)

type usb interface {
	Start()
	Stop()
}

type phone struct {
	Name string
}

func (p phone) Start() {
	fmt.Println("手机开始工作")
}

func (p phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p phone) Call() {
	fmt.Println("手机在打电话")
}

type camera struct {
	Name string
}

func (c camera) Start() {
	fmt.Println("相机开始工作")
}

func (c camera) Stop() {
	fmt.Println("相机停止工作")
}

type computer struct {
}

func (com computer) Working(u usb) {
	u.Start()
	if phone, ok := u.(phone); ok {
		phone.Call()
	}
	u.Stop()
}

func main() {
	var usbArr [3]usb
	usbArr[0] = phone{"vivo"}
	usbArr[1] = phone{"小米"}
	usbArr[2] = camera{"尼康"}

	var com computer
	for _, v := range usbArr {
		com.Working(v)
		fmt.Println()
	}
}
