package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n

}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25
	// * is the value of a memory address & is the memory address of a value
	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i:", i)

	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
