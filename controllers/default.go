package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Email"] = "xbaofeng@gmail.com"
	c.TplNames = "index.tpl"
}
