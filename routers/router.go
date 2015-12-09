package routers

import (
	"github.com/xuebaofeng/stock-web-golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
