// Made by: Mohamad
// Date: Mar 23 2023
//
// This script calculates the area of an icosagon.

package main

import (
	"fmt"
	"math"
)

func main() {
	// declare variables
	var sideLength float64
	var sides float64 = 20

	// get user input
	fmt.Println("Enter the length of the side of the icosagon: ")
	fmt.Scanln(&sideLength)

	// calculate area
	area := (float64(sides) * math.Pow(sideLength, 2)) / (4 * math.Tan(math.Pi/float64(sides)))

	// Output result
	fmt.Println("The area of the icosagon is: ", area, "cm²")

	fmt.Println("\nDone.")
}
