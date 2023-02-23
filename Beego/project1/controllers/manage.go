package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"project1/models"
)

type MngController struct {
	beego.Controller
}

func (manage *MngController) Add() {
	manage.Data["Form"] = &models.Article{}
	//manage.Layout = "basic-layout.tpl"
	manage.LayoutSections = make(map[string]string)

	manage.TplName = "add.tpl"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}

	if err := manage.ParseForm(&article); err != nil {
		logs.Error("Couldn't parse the form. Reason: ", err)
	} else {
		manage.Data["Articles"] = article
		valid := validation.Validation{}
		isValid, _ := valid.Valid(article)

		if manage.Ctx.Input.Method() == "POST" {
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				logs.Error("Form didn't validate.")
			} else {
				//searchArticle := models.Article{Name: article.Name}
				logs.Debug("Article name supplied:", article.Name)
				//err = o.Read(&searchArticle)
				logs.Debug("Err:", err)
				id, err := o.Insert(&article)
				if err == nil {
					msg := fmt.Sprintf("Article inserted with id:", id)
					logs.Debug(msg)
					flash.Notice(msg)
					flash.Store(&manage.Controller)
					fmt.Println("asdasd")
					manage.Redirect("http://localhost:8080/", 302)
				} else {
					msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
					logs.Debug(msg)
					flash.Error(msg)
					flash.Store(&manage.Controller)
				}
			}
		}
	}
}
