package main

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	"project1/models"
	_ "project1/routers"
)

func main() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=demo sslmode=disable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connection successfyull")
	}
	orm.RunSyncdb("default", false, true)
	orm.RegisterModel(new(models.Article))
	beego.Run()
}
