package main

import "fmt"

func main() {
	var names = [...]string{"xyz", "dasf", "sda", "Awawe"}
	names1 := [10]string{}

	names1[0] = "1"
	names1[1] = "2"
	//fmt.Print(names, names1)
	fmt.Printf("%s\n", names)
	fmt.Printf("%#v", names)
	fmt.Printf("\n%q", names)

}
