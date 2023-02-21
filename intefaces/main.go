package main

import "fmt"

type Building interface {
	area()
	volume()
}

type Values struct {
	height int
	width  int
	length int
}

func (v Values) area() {
	ans := v.width * v.length
	fmt.Println("Area of Building is:", ans)
}

func (v Values) volume() {
	ans := v.width * v.height * v.length
	fmt.Println("Volume of building is:", ans)
}

func main() {
	//fmt.Println("here")
	var b Building = Values{10, 5, 7}
	b.volume()
	b.area()
}
