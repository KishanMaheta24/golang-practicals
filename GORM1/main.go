package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"models/models"
	"os"
)

// type employee=
var db *gorm.DB
var err error

func main() {
	dsn := "kishan:Simform@123@tcp(127.0.0.1:3306)/new_schema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database")
		os.Exit(1)

	} else {
		fmt.Println("successfully connected", db)
	}
	//db.Take("employee").AutoMigrate(&models.Employee{})
	//fetchEmployee()
	//fetchBranch()
	//fetchEmployeeDateOrderBySN()
	//findBornAfter1969()
	//inCondition()
	//grpBy()
	//subqueryExample()
	//joinsExamples()
}

func fetchEmployee() {
	var employee []models.Employee
	db.Table("employee").Find(&employee)
	//fmt.Println(employee[0].Birth_date)
	pretty(employee)

}

func fetchBranch() {
	var branch []models.Branch
	db.Table("branch").Find(&branch)
	//pretty(employee)

}

// Find all employees ordered by sex then name
func fetchEmployeeDateOrderBySN() {
	var employee []models.Employee
	db.Table("employee").Order("sex").Find(&employee)
	pretty(employee)

}

// Find all employees who are female & born after 1969 or who make over 80000
func findBornAfter1969() {
	var employee []models.Employee
	db.Table("employee").Where("salary>80000 AND birth_date>=1970-01-01").Find(&employee)
	pretty(employee)
}

func inCondition() {
	var employee []models.Employee
	db.Table("employee").Where([]int64{100, 102, 103}).Find(&employee)
	pretty(employee)
}

func pretty(employee []models.Employee) {
	js, err := json.MarshalIndent(employee, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(js))
	}
}

func grpBy() {
	var work []models.Works_With
	db.Table("works_with").Select("client_id,sum(total_sales) as Total_sales").Group("client_id").Find(&work)
	js, err := json.MarshalIndent(work, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(js))
	}
}

func subqueryExample() {
	var works []models.Works_With

	db.Table("works_with").Where("total_sales > (?)", db.Table("works_with").Select("AVG(total_sales)")).Find(&works)

	js, err := json.MarshalIndent(works, "", "\t")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(js))
	}
}

func joinsExamples() {
	type temp struct {
		Emp_id      int
		First_name  string
		Branch_name string
	}
	var tmp []temp
	db.Table("employee").Select("emp_id,first_name,branch.branch_name as branch_name").Joins("inner join branch on employee.emp_id=branch.mgr_id").Find(&tmp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tmp)
	}
}
