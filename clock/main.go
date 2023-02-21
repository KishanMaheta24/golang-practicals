package main

import (
	"fmt"
	time2 "time"
)

func main() {
	type numbers [5]string

	zero := numbers{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := numbers{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := numbers{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := numbers{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := numbers{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := numbers{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := numbers{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := numbers{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := numbers{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := numbers{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]numbers{zero, one, two, three, four, five, six, seven, eight, nine}
	for {
		time := time2.Now()
		//fmt.Println(time.Hour(), time.Minute(), time.Second())
		time_arr := [...]numbers{digits[time.Hour()/10], digits[time.Hour()%10], digits[time.Minute()/10], digits[time.Minute()%10], digits[time.Second()/10], digits[time.Second()%10]}
		for i, _ := range time_arr[0] {
			for j, _ := range time_arr {
				fmt.Print(time_arr[j][i], "    ")
				//fmt.Println(i, j)
			}
			fmt.Println()
		}
		time2.Sleep(time2.Second)
		fmt.Println(" ")

	}
}
