package main

import "fmt"

const (
	a = iota //iota=0 so a=0...after this every statement in const will get incremented itself whether iota keyword is used or not
	b        // b gets incremented to 1
	c = iota //c gets incremented to 2,presence of iota keyword does not affect its value
)
const ( //iota get redeclared when const keyword is again used
	d = iota
	e = 33//here iota increments and becomes 1 but is not used due to overriding value of e->33
	f //f will take value of e (i.e 33)...Still iota increments and becomes 2
	g// iota becomes 3 but still is not used as g overrides f=33 hence g=33
	h=iota//h will take incremented value of iota , (i.e) h=4
	i//iota increments to 5,hence i=5
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}

