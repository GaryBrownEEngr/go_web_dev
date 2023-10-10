package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Let's Practice Addition.")
	startTime := time.Now()
	correctCount := 0
	totalCount := 0

	for {
		a := getARandomNumber(0, 12)
		b := getARandomNumber(0, 12)
		answer := a + b

		seconds := time.Since(startTime).Seconds()
		fmt.Printf("%d+%d=?", a, b)
		fmt.Printf("                                                                  ")
		fmt.Printf("[%d/%d", correctCount, totalCount)
		fmt.Printf(":%.0f%%:%.1fs", 100.0*float64(correctCount)/float64(totalCount), seconds)
		fmt.Printf(":%.1fs/p]\n", seconds/float64(correctCount))
		totalCount++
		usersAnswer, err := getNumberFromConsole()
		if err != nil {
			fmt.Println("oops! you didn't give a number")
			fmt.Println("The answer is: ", answer)
			continue
		}

		if answer == usersAnswer {
			fmt.Printf("great!\n\n")
			correctCount++
		} else {
			fmt.Printf("try again. The answer is %d\n\n", answer)
		}
	}
}

func getNumberFromConsole() (int, error) {
	var xStr string
	fmt.Scanln(&xStr)

	x, err := strconv.Atoi(xStr)
	return x, err
}

// [min, max]
func getARandomNumber(min, max int) int {
	randRange := max - min + 1
	return rand.Intn(randRange) + min
}
