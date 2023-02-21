package main

import (
	"fmt"
	time2 "time"
)

func DoingNothing(inp string) {
	for i := 0; i < 10; i++ {
		fmt.Println(inp, ":", i)
		time2.Sleep(time2.Millisecond * 200)
	}
}
func DoingNothing1(inp string) {
	for i := 0; i < 10; i++ {
		fmt.Println(inp, ":", i)
		time2.Sleep(time2.Millisecond * 200)
	}
}

func main() {
	go DoingNothing("1")
	//time2.Sleep(time2.Millisecond * 200)
	go DoingNothing1("2")
	DoingNothing("")
	//time2.Sleep(time2.Millisecond * 200)

}
