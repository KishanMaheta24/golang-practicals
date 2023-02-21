package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mymodules/models"
	"os"
)

type Users = models.Users

var db *gorm.DB
var err error
var User Users

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "kishan:Simform@123@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error connecting to database")
		os.Exit(1)

	} else {
		fmt.Println("successfully connected", db)
	}
	//findUserTake()
	//findUserLast()
	//findById(78)
	//findByCondition()
	//whereClause()
	//whereInClause()
	//whereMultiConditions()
	//orderBy()
	//selectSpecificFields()
	limitOffset()
}

func createUser() {
	//insert into users values (77,kishan,asdaasf)
	createUser := Users{77, "kishan", "asdjahd"}
	result := db.Create(&createUser)
	fmt.Println(result)

}

func createUsingOmit() {
	//omit function omits the columns given as parameters
	user := Users{Name: "xyz", Email: "afdja@gmail.com"}
	result := db.Omit("Id").Create(&user)
	fmt.Println(result)
}

func createBatchInsert() {
	users := []Users{
		{79, "xyz1", "xyz1@gmail.com"},
		{80, "xyz2", "xyz2@gmail.com"},
		{81, "xyz3", "xyz3@gmail.com"}}

	result := db.Create(&users)
	fmt.Println(result)
	fmt.Println(users)
}

func findUserFirst(db *gorm.DB) {
	//SELECT * FROM users ORDER BY id LIMIT 1
	var user Users
	db.Find(&user)
	fmt.Println(user)
}
func findUserTake() {
	// SELECT * FROM users LIMIT 1;
	var user Users
	db.Take(&user)
	fmt.Println(user)
}

func findUserLast() {
	// SELECT * FROM users LIMIT 1;
	var user Users
	res := db.Last(&user)
	fmt.Println(user, res)
}

func findById(id int) {
	var user = Users{Id: id}
	res := db.First(&user)
	fmt.Println(res, user)
}

func findByCondition() {
	var User []Users
	db.Where("name LIKE ?", "xyz%").Find(&User)
	fmt.Print(User)
}
func whereClause() {
	//var User Users
	db.Where("name=?", "xyz2").Find(&User)
	fmt.Println(User)
}

func whereInClause() {
	db.Where([]int{77, 79, 80}).Find(&User)
	fmt.Println(User)
}

func whereMultiConditions() {
	var users []Users
	db.Where(map[string]interface{}{"name": "kishan", "id": "77"}).Find(&users)
	fmt.Println(users)
}

func orderBy() {
	var users []Users
	db.Order("id desc").Find(&users)
	fmt.Println(users)

}

func selectSpecificFields() {
	var users []Users
	db.Select("id", "name").Find(&users)
	fmt.Println(users)
}

func limitOffset() {
	var users []Users
	db.Limit(2).Offset(1).Order("id desc").Find(&users)
	fmt.Println(users)
}
