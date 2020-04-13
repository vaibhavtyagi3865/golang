package main

import "fmt"

type custom_type int

var e custom_type = 221
var f int

func main() {
	fmt.Printf("e=%v\n", e)
	fmt.Printf("Type(e)=%T\n", e)
	e = 42
	fmt.Printf("e=%v\n", e)
	f = int(e)
	fmt.Println("f=", f)
}
