package controllers

import (
	"GORM/models"
	"GORM/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	db, err := utils.InitConnection()
	if err != nil {
		fmt.Println("error")
	} else {
		db.Create(&models.Item{Id: 75, Code: "D44", Price: 100})
	}

}
