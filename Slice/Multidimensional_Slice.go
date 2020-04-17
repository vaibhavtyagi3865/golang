package main

import "fmt"

func main() {
	//1st method
	x := [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {1, 3, 5, 7}} //also written as x := [3][4]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {1, 3, 5, 7}}
	fmt.Println("x[][]:", x)
	fmt.Println("x[0][2]:", x[0][2])
	//2nd way we can also use oversized array and assign less values to it
	//empty places are occupied with '0'
	x2 := [3][5]int{{11, 22, 33, 44}, {12, 14, 16, 18}, {10, 13, 15, 17}} //3 arrays with every subarray containing 5 elements
	fmt.Println("x2[][]:", x2)

	//3rd method: using make method
	y := make([][]int, 3) //y= an array which stores 3 arrays of type int
	for i := 0; i < 3; i++ {
		y[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			y[i][j] = i + j
		}
	}
	fmt.Println("y[][]:", y)

	//4th method: Using copy method
	z := make([][]int, 3)
	for i := 0; i < 3; i++ {
		z[i] = make([]int, len(x[i]))
		copy(z[i], x[i])
	}
	fmt.Println("copied array z[][]:", z)

	//5th method-----------------------------------------------------
	var value int
	a := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 0, 4) //Here we have taken length as 0 otherwise the command:-
		// "a[i]=make([]int,4)" will create a[i] as [0,0,0,0] and other values will be appended after 0's as
		// [0,0,0,0,1,2,and so on ]. To avoid this we define length as 0 and capacity as 4 so a[i]=[empty slice]
		for j := 0; j < 4; j++ {
			value = (2 * i) + j
			a[i] = append(a[i], value)
		}

	}
	//fmt.Println("a[][]",a)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	//appending in 2d slice
	arr := make([][]int, 4)
	for i := 0; i < 3; i++ {
		arr[i] = make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			value = (2 * i) + j
			arr[i] = append(arr[i], value)
		}

	}
	fmt.Println("1st arr[][]:", arr)
	duplicate := []int{100, 101, 102, 103}
	arr[3] = append(arr[3], duplicate...)
	fmt.Println("arr[][] after appending a 1d array inside it:", arr)

	//Original
	//arr = [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {1, 3, 5, 7}}
	//fmt.Println("Original arr[][]:",arr)
	//fmt.Println()
	//duplicate:=[]int{100,101,102,103}
	//arr[3]=append(arr[3] ,duplicate...) //why this doesn't work??
	//fmt.Println("arr[][] after appending a 1d array inside it:",arr)
}
