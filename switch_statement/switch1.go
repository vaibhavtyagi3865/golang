package main

import "fmt"

func main() {
	var x = 4
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 4:
		fmt.Println("Four")
		fallthrough //next statement will always run after this keyword,whether right or wrong
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	default: //works only if no above case is true, also works with 'fallthrough'
		fmt.Println("Default Output")
	}
}
