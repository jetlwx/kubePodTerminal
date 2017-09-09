package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jetlwx/kubePodTerminal/routers"
)

func main() {
	beego.Run()
}
