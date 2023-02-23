package models

type User_details struct {
	Id        int    `form:"-"`
	Name      string `form:"name,text,name"`
	Email     string `form:"email,email,email"`
	Password  string `form:"password,password,password"`
	Contact   int    `form:"contact,number,contact"valid:"MinSize(10);MaxSize(10)"`
	TechStack string `form:"stack,text,stack"`
}

func (u *User_details) TableName() string {
	return "user_details"
}
