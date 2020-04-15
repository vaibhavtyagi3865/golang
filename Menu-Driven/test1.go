package main

import "fmt"

var val int
var x [5]int

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Enter value",(i+1) )
		fmt.Scanln(&val)
		x[i] = val
	}
	fmt.Println(x)

}
