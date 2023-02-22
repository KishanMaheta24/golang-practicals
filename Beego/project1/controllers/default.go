package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) AboutUs() {
	main.Data["Name"] = "Kishan Maheta"
	main.Data["Email"] = "kishan.m@simformsolutions.com"
	main.Data["Developer"] = "Golang Developer"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "aboutus.tpl"
}
