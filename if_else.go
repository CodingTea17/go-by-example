package main

import "fmt"

func main() {

	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("8 is divisivle by 4")
	}

	if n := 9; n < 9 {
		fmt.Println(n, "is a neg num")
	} else if n < 10 {
		fmt.Println(n, "has one 1 digit")
	} else {
		fmt.Println(n, "has multiple digits")
	}
}
