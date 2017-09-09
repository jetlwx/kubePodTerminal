package routers

import (
	"github.com/astaxie/beego"
	"github.com/jetlwx/kubePodTerminal/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pod/sub", &controllers.MainController{}, "post:Sub")
}
