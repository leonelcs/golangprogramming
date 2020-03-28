package main

import (
	"fmt"
)

func main() {
	fmt.Println("d1 execution")
	d1()
	fmt.Println("d2 execution")
	d2()
	fmt.Println("d3 execution")
	d3()
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	defer fmt.Println()
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	defer fmt.Println()
}
