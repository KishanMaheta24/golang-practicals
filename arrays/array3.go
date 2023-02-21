package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

var escapingWords = [...]string{"and", "or", "was", "the", "since", "very"}

func main() {

	corpus := strings.Split(corpus, " ")
	fmt.Println(corpus)
	args := os.Args[1:]
forloop:
	for i, obj := range corpus {

		for _, escwords := range escapingWords {
			if escwords == obj {
				continue forloop
			}
		}

		for _, words := range args {
			//fmt.Print(obj, words, "w\n")
			if words == obj {
				fmt.Println("#", i, obj)
			}
		}
	}
}
