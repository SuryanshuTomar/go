package main

import "fmt"

func main() {
	// Go have only one type of loop and that is "for" loop. Go doesn't have the "while" loop.

	// Method 1 - to use the "for" loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Method 2 - to use the "for" loop - use it as a while loop
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("Sum is : ", sum)

	// Method 3 - to use the "for" loop - if you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	// for {
	// 	fmt.Println("Running...Infinitely...")
	// }
}
