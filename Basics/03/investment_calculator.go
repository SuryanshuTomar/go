package main

import (
	"fmt"
	"math"
)

// Approach 1
// func main() {
// 	var investmentAmount = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10

// 	var futurevalue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

// 	fmt.Print("Future value of ", investmentAmount, " investment with an expected return rate of ", expectedReturnRate, " in ", years, "years will be : ", futurevalue)
// }

// Approach 2
func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	futurevalue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futurevalue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println("Future value of ", investmentAmount, " investment with an expected return rate of ", expectedReturnRate, " in ", years, "years will be : ", futurevalue)
	fmt.Println("And Real Future value with inflation will be ", futureRealValue)
}