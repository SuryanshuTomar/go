package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println("Conditionals : ")

	var num int64
	fmt.Print("Enter a number : ")
	fmt.Scanln(&num)

	// Method 1 : If - else if - else
	if num < 50 {
		fmt.Println("You number is lesser than 50")
	} else if num == 50 {
		fmt.Println("You number is 50")
	} else {
		fmt.Println("You number is greater than 50")
	}

	// Method 2 : The if statement can start with a short statement to execute before the condition. 
	// Variables declared by the statement are only in scope until the end of the if.
	// Variables declared inside an if short statement are also available inside any of the else blocks.
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
