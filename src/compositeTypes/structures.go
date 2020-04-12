package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	var s1 aStructure
	p1 := aStructure{"fmt", 12, -2}

	s1.height = 10
	s1.person = "abc"
	s1.weight = 10

	fmt.Println("s1: ", s1)
	fmt.Println("p1: ", p1)
	p2 := aStructure{weight: 12, height: -2}
	fmt.Println("s2: ", p2)

	var v aStructure
	fmt.Println("V.person: ", v.person, " V.height: ", v.height, " V.weight: ", v.weight)

}
