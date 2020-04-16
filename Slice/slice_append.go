package main

import "fmt"

func main() {
	//1st way
	x := []int{221, 21, 5, 2, 1}
	x = append(x, 24, 48, 72)
	fmt.Println("x[]:", x)
	//2nd way
	y := []int{2, 4, 6, 8, 9}
	x = append(x, y...) // '...' is used to take out all values from slice y and insert them in slice x
	fmt.Println("x[]:", x)
}
