package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d ", i)
	}

	//there is no while, so use the for
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Printf(" %d ", i)
		i++
	}
	i = 0
	for ok := true; ok; ok = i < 10 {
		fmt.Printf(" %d ", i)
		i++
	}

	anArray := [5]int{0, 1, -1, 2, -2}
	for j, value := range anArray {
		fmt.Println("index:", j, "value: ", value)
	}

}
