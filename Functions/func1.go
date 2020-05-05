package main

import (
	"fmt"
)

func add(x int, y int) int { //Function Returning Single Value
	z := x + y
	return z
}

func calculate(l int, b int) (area int, parameter int) { //function returning multiple values
	area = l * b
	parameter = 2 * (l + b)
	return //Simply return the statement without specifying variable name
}

func main() {
	a := 6
	b := 7
	fmt.Println(add(a, b))

	var length int
	var breadth int
	fmt.Println("Enter Length")
	fmt.Scanln(&length)
	fmt.Println("Enter Breadth")
	fmt.Scanln(&breadth)

	area, parameter := calculate(length, breadth) //This function returns more than 1 value(i.e area and parameter)
	fmt.Println("Area:", area)
	fmt.Println("Parameter:", parameter)
}

//--------------------------------------------------------------
