package main

import (
	"fmt"
	"time"

	"code.logagent/conf"

	"code.logagent/kafka"
	"code.logagent/taillog"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	//1，读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2, 发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Load ini failed, err:%v\n", err)
		return
	}
	//初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("Init kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success!")
	//打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success!")
	run()
}
