package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	loop()
}
func loop() {
	fmt.Println("This is a Go programming language tutorial")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
