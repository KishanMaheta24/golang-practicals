package routers

import (
	"beegoProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "post,get:Add")
}
