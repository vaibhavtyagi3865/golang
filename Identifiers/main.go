package main

import "fmt"

var str = "hello this is me"

var z = 5 //For global scope we use var keyword
var g int //if global variable is not assigned any value ,then we have to mention its type,otherwise we will get error
func main() {
	x := 40 // ":=" will not work outside the local scope .
	fmt.Println("x=", x)
	x = 5
	fmt.Println("x again=", x)
	fmt.Println("z=", z)
	foo()
}
func foo() {
	fmt.Println("z again=", z)
	fmt.Printf("Data type of variable z=%T\n", z)
	fmt.Println("g=", g)
	fmt.Println(str)
}
