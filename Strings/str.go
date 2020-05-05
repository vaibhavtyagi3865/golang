package main

import "fmt"

var str string
var f uint=7
func main() {
	str="Hello World"
	ar:=[]uint8(str)//type conversion
	fmt.Println(len(ar))

	for i:=0;i<len(ar);i++{
		fmt.Printf("%c %v\n",str[i],ar[i])
	}

	//for i,v:=range str{
	//	fmt.Println(i,v)
	//}
}

