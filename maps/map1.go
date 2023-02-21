package main

import (
	"fmt"
	"os"
	"strings"
)

type UserDetails struct {
	Name       string
	Email      string
	Favourites map[string]string
}

func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Print("enter name")
	} else if len(args) == 2 {
		fmt.Println("enter email")
	} else if len(args) == 3 {
		fmt.Println("enter atleast one favourite item")
	} else {
		fvrts := make(map[string]string)
		for _, value := range args[3:] {
			fmt.Println(value)
			fvrts[strings.Split(value, ":")[0]] = strings.Split(value, ":")[1]
		}

		user1 := UserDetails{args[1], args[2], fvrts}
		fmt.Println(user1)
		fmt.Println(user1.getName())
	}

}

func (u UserDetails) getName() string {
	return u.Name
}

func (u UserDetails) getEmail() {
	fmt.Print(u.Email)
}
