
package main

import "math"

func getRandomNumber(min int, max int) int {
	return min + (max-min)*rand.Float64()
}

func main() {
	// Use the above function to generate a random number between 1 and 10
	fmt.Println(getRandomNumber(1, 10))
}