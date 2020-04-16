package main

import "fmt"

func main() {
	var x = [5]int{1, 2, 3}
	fmt.Println("x[]:", x)
	fmt.Println("len(x):", len(x))
	fmt.Println("cap(x):", cap(x))
	//creating slice with make method
	// syntax:-  make(type of slice,length of slice,capacity of slice)--->capacity could be optional
	y := make([]int, 3, 7)

	//Slice ,even if it is empty will show a slice of length(here 3) as defined above
	//and will display a slice containing 0's. Since length of slice y=3,hence y=[0,0,0]
	fmt.Println("empty y[]:", y)

	fmt.Println("len(y):", len(y))
	fmt.Println("cap(y):", cap(y))

	y = append(y, 4, 5, 6)
	fmt.Println("After append operation 1, y[]:", y)
	fmt.Println("len(y):", len(y))
	fmt.Println("cap(y):", cap(y))

	y = append(y, 10)
	fmt.Println("After append operation-2, y[]:", y)
	fmt.Println("len(y):", len(y))
	fmt.Println("cap(y):", cap(y))
	//Till this point number of elements in slice y equals capacity of slice y,
	//Further, if we add another element in 'y', then the capacity of slice 'y' will double itself
	//Also length of slice y will also change itself
	y = append(y, 11)
	fmt.Println("After append operation-3, y[]:", y)
	fmt.Println("new len(y):", len(y))
	fmt.Println("new cap(y):", cap(y))
}
