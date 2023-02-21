package main

import "fmt"

func main() {
	sci_info := [...][3]string{
		{"albert", "einstein", "time"},
		{"isaac", "newton", "apple"},
		{"stephen", "hawkins", "blackhole"},
		{"marie", "curie", "radium"},
		{"charles", "darwin", "fittest"}}

	fmt.Println("FirstName  lastname  nickname")

	for i := 0; i < len(sci_info); i++ {
		for j := 0; j < len(sci_info[0]); j++ {
			fmt.Printf("%-11s", sci_info[i][j])
		}
		fmt.Println("")
	}
}
