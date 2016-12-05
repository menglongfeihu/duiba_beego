package main

import (
	_ "duiba_activity/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogFuncCall(true)
	beego.Info("========开始启动程序========")
	beego.Run()
}
