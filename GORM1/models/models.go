package models

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	Emp_id     int "gorm:primaryKey"
	First_name string
	Last_name  string
	Birth_date time.Time "json:birth_date"
	Sex        string
	Salary     int
	Super_id   int
	Branch_id  int
}

type Branch struct {
	Branch_id      int "gorm:primaryKey"
	Branch_name    string
	Mgr_id         int
	Mgr_start_date time.Time
}
type Client struct {
	gorm.Model
	Client_id   int "gorm:primaryKey"
	Client_name string
	Branch_id   int
}
type Works_With struct {
	Emp_id      int "gorm:primaryKey"
	Client_id   int "gorm:primaryKey"
	Total_sales int
}

type Branch_suppliers struct {
	//gorm.Model
	Branch_id     int    "gorm:primaryKey"
	Supplier_name string "gorm:primaryKey"
	Supply_type   string
}
