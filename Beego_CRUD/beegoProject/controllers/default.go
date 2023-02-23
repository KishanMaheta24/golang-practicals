package controllers

import (
	"beegoProject/models"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Add() {
	c.Data["Form"] = &models.User_details{}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	flash := beego.ReadFromRequest(&c.Controller)

	if ok := flash.Data["error"]; ok != "" {
		c.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	user_details := models.User_details{}

	if err := c.ParseForm(&user_details); err != nil {
		logs.Error("form parsing failed...", err)
	} else {
		c.Data["user_details"] = user_details
		valid := validation.Validation{}
		isValid, _ := valid.Valid(user_details)

		if c.Ctx.Input.Method() == "POST" {
			if !isValid {
				c.Data["Errors"] = valid.ErrorsMap
				logs.Error("Form didn't validate.")
			} else {
				id, err := o.Insert(&user_details)
				if err == nil {
					msg := fmt.Sprintf("inserted id:", id)
					flash.Notice(msg)
					flash.Store(&c.Controller)
				} else {
					msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
					logs.Debug(msg)
					flash.Error(msg)
					flash.Store(&c.Controller)
				}
			}
		}

	}

}
