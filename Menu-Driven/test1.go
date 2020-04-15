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
	//2nd method
	for i:=0;i<5;i++{
		fmt.Println(x[i])
	}
	//3rd method
	for i, v:=range x{   //here i=index and v=value for this type of loop
		fmt.Println("Index:",i,"Value:",v)
	}

}
