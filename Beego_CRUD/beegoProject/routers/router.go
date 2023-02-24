package routers

import (
	"beegoProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "post,get:Add")
	beego.Router("/search", &controllers.MainController{}, "post,get:Search")
	beego.Router("/delete", &controllers.MainController{}, "post,get:Delete")

}
