package main

import (
	"beegoProject/models"
	_ "beegoProject/routers"
	"fmt"
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("driver successful")
	}
	err = orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=demo sslmode=disable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connection successful")
	}
	orm.RegisterModel(new(models.User_details))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sync successful")
	}

}

func main() {

	beego.Run()
}
