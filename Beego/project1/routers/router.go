package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"project1/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/aboutUs/:id([0-9]+)", &controllers.MainController{}, "get:AboutUs")
}
