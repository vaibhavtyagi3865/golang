package main

import "fmt"

func main() {
	fmt.Println("Println function can accept unlimited no. of arguments, like int-", 221, "boolean-", true, "and strings")

	n, err := fmt.Println("\n1st function")             //return type of Println function is ===>(no of bytes,error)
	fmt.Println("no of bytes in above statement:", n) //n= no. of bytes in fmt function
	fmt.Println("error in above function:", err)      //err=tells if any error present in fmt function

	_, e := fmt.Println("\n2nd function") // if any variable except _ symbol is a new variable on LHS of = expression then we will use := between LHS and RHS
	fmt.Println("error=", e)            //otherwise we will use = symbol between LHS and RHS

	n, _ = fmt.Println("\n3rd function") //m variable is reused here so this time we will use '=' instead of ':=' symbol
	fmt.Println("no of bytes in above statement:", n)

}
