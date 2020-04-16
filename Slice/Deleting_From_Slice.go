package main

import "fmt"

func main() {
	x := []int{221, 21, 5, 2, 1, 44, 55, 66, 77, 88, 99}
	fmt.Println("x[1:4]:", x[1:4]) //using this we can display elements from index 1 to 3 (4 is exclusive)
	var y []int
	y = append(x[0:4], x[8:9]...) // gives [221,21,5,2,77]
	fmt.Println("y[]:", y)
	fmt.Println("new x[]:", x)
	//var z []int
	//z=append(x[0:5],x[9])
	//fmt.Println("z[]:",z)
	//fmt.Print(append(x[0:5],x[9]))
}
