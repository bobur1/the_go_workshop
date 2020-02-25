package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 101; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}

		if i%3 == 0 {
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
		}

		//in order to print the next number or string in new line
		fmt.Print("\n")
	}
}
