//2nd way of running switch commands
package main

import "fmt"

func main() {
	 n:= "Tyagi"
	switch {

	case n == "Verma", n == "Kumar", n == "Taneja": //we can take multiple conditions in a case
		fmt.Println("Verma or Kumar or Taneja")
	case n == "Sharma":
		fmt.Println("Sharma")
	case n == "Tyagi":
		fmt.Println("Tyagi")
	case n == "Jindal":
		fmt.Println("Jindal")
		fallthrough
	default:
		fmt.Println("Gupta")
	}
}
