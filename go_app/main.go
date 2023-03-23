// Made by: Mohamad
// Date: Mar 23 2023
//
// This script calculates the area of an isosagon.

package main

import (
	"fmt"
	"math"
)

func main() {
	// declare variables
	var sideLength float64
	sides := 20

	// get user input
	fmt.Println("Enter the length of the side of the isosagon: ")
	fmt.Scanln(&sideLength)

	// calculate area
	area := (float64(sides) * math.Pow(sideLength, 2)) / (4 * math.Tan(math.Pi/float64(sides)))

	// Output result
	fmt.Println("The area of the isosagon is: ", area, "cm²")

	fmt.Println("Done.")
}
