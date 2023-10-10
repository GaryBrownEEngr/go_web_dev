package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")

	// dogs := 2
	// fmt.Printf("i love %d dogs!\n", dogs)

	fmt.Println("How many lizards do you love?")

	var xStr string
	fmt.Scanln(&xStr)

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic("oops! you didn't give a number")
	}

	if x > 10 {
		fmt.Println("WOW! That's a lot of picking up poo to do.")
	}

	if x < 0 {
		fmt.Println("WHAT IS A NEGATIVE lizard???.")
		return
	}

	if x == 1 {
		fmt.Printf("You like %d lizard too!\n", x)
	} else {
		fmt.Printf("You like %d lizards too!\n", x)
	}
}
