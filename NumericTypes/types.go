package main

import (
	"fmt"
	"runtime"
)

var x1 uint8  //range 0 to 255
var x2 int8   //range -128 to 127
var x3 uint16 //range 0 to 65535
func main() {
	//x1 = -5 //will give error
	x1 = 5
	fmt.Println("x1=", x1)

	//x2=-180 // gives error because of range of x2
	x2 = -120 //works fine
	fmt.Println("x2=", x2)
	//x3 = 70000 //gives error,not in range of uint16
	x3 = 65530
	fmt.Println("x3=", x3)

	//To print Go Architectur and Go OS,we import 'runtime' package
	fmt.Println("Go Operating System:", runtime.GOOS)
	fmt.Println("Go Runtime Architecture:", runtime.GOARCH)
}
