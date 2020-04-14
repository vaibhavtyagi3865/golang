package main

import (
	"fmt"
)

const (
	_  = iota             //iota=0,no variable thrown to use this value
	kb = 1 << (iota * 10) //iota increments to 1 and 1<<10 is equal to 2^10=1024
	mb = 1 << (iota * 10) //iota increments to 2 and 1<<20 is equal to 2^20=2048
	gb = 1 << (iota * 10) //iota increments to 3 and 1<<30 is equal to 2^30=4096

)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
