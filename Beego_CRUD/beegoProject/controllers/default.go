package controllers

import (
	"beegoProject/models"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
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

func (c *MainController) Search() {
	c.TplName = "search.tpl"

	flash := beego.ReadFromRequest(&c.Controller)
	if ok := flash.Data["error"]; ok != "" {
		c.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	if c.Ctx.Input.Method() == "POST" {
		name := c.GetString("name")
		fmt.Println("name:", name)

		if name == "" {
			flash.Notice("please enter name")
			flash.Store(&c.Controller)
		} else {
			res := models.User_details{}
			qs := o.QueryTable("user_details")
			_, err := qs.Filter("name", name).All(&res)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				//res, err := json.Marshal(res)
				if err != nil {

					flash.Notice("error")
					flash.Store(&c.Controller)
				} else {
					c.Data["user"] = res

					fmt.Println(res)
				}

			}
		}
	}

}

func (c *MainController) Delete() {
	c.TplName = "delete.tpl"

	flash := beego.ReadFromRequest(&c.Controller)
	if ok := flash.Data["error"]; ok != "" {
		c.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	if c.Ctx.Input.Method() == "POST" {
		id := c.GetString("name")
		fmt.Println("name:", id)

		if id == "" {
			flash.Notice("please enter ID")
			flash.Store(&c.Controller)
		} else {
			res := models.User_details{}
			temp, err := strconv.Atoi(id)
			res.Id = temp
			resp, err := o.Delete(&res)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				//res, err := json.Marshal(res)
				if err != nil {

					flash.Notice("error")
					flash.Store(&c.Controller)
				} else {
					c.Data["user"] = "record deleted successfully" + string(resp)
					fmt.Println(res)
				}

			}
		}
	}

}
