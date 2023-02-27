package main

import (
	_ "GORM/routers"
	"GORM/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	// Migrate the schema
	//Db.AutoMigrate(&models.Item{})
	_, err := utils.InitConnection()

	if err != nil {
		fmt.Println("error while connecting to database")
	}
	beego.Run()
}
