package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Println(s2)

	integer := make([]int, 20)
	for i := 0; i < 20; i++ {
		fmt.Print(" ", integer[i], " ")
	}
	fmt.Println()

	integer = append(integer, 12345)
	for i, value := range integer {
		fmt.Print("index ", i, " value ", value, " ")
	}
	fmt.Println()

	s3 := integer[10:20]

	for i, value := range s3 {
		fmt.Print("index ", i, " value ", value, " ")
	}
	fmt.Println()

	fmt.Println("length: ", len(integer), " capacity ", cap(integer))
}
