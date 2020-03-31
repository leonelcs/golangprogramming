package main

import (
	"fmt"
)

func main() {
	//slice behavior
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integers := make([]int, 2)
	fmt.Println(integers)
	integers = nil
	fmt.Println(integers)

	//use of [x:y]
	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:] //keeps the reference

	fmt.Println(anArray)
	fmt.Println(refAnArray)
	anArray[4] = -100
	fmt.Println(refAnArray)
	//using make to create 1 and 2 dimension slices
	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	//manual initialization
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}
	fmt.Println(twoD)
	fmt.Println()

	//using range
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
		fmt.Println()
	}

}
