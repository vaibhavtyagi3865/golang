package main

import "fmt"

var a int = 42
var b = "James Bond"
var c = true

func main() {
	s := fmt.Sprintf("a=%v,b=%v,c=%v", a, b, c)
	fmt.Println(s)
}
