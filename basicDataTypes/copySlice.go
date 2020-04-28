package main

import (
	"fmt"
)

func main() {
	array6 := []int{1, 2, 3, 4, 5, 6}
	array4 := make([]int, 4)

	//copy will respect the destination size
	//copy(dest, src)
	copy(array4, array6)

	fmt.Println(array6)
	fmt.Println(array4)
}
