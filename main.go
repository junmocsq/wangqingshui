package main

import (
	"github.com/astaxie/beego"

	_ "github.com/astaxie/beego/session/redis"
	_ "wangqingshui/routers"
)

func main() {

	// beego 热启动 kill -HUP 进程id
	//beego.BConfig.Listen.Graceful = true // 有问题暂不启动
	beego.Run()
}
