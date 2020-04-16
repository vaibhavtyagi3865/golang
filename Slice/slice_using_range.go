package main

import (
	"fmt"
)
//1st method of making slice
var x =[8]int{1, 2, 3, 4, 5}

func main() {
	//1st method to display slice
	fmt.Println("x[]:",x)
	//2nd method to display slice
	fmt.Println("index Value")
	for i, v := range x {//here i=index ,v=value of elements in loop
		fmt.Println(i,"   ", v)
	}

	//2nd method of making a slice
	y:=[5]int{221,21,5,2,1}
	fmt.Println("y[]:",y)

	//3rd method of making slice
	var z=[4]int{1,2,3}
	fmt.Println("z[]:",z)

}
