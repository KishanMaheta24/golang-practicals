package controllers

import (
	"Example/database"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post() {
	var req database.UserForm
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	o := orm.NewOrm()
	o.Using("default")
	if _, err := o.Insert(&req); err != nil {
		this.Ctx.Output.SetStatus(309)
		this.Data["json"] = "Email already used"
		this.ServeJSON()
		return
	}
	this.Ctx.Output.SetStatus(200)
	//dec := json.NewDecoder(this.Ctx.Request.Body)
	//err := dec.Decode(&req)
	//if err == nil {

	this.Data["json"] = req
	this.ServeJSON()

	//}
	return
}
