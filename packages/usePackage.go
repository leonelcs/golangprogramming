package main

import (
	"fmt"

	"github.com/leonelcs/golangprogramming/packages/apackage"
)

func main() {
	fmt.Println("esse é a main")
	apackage.PrintPackage()

	fmt.Println("Myconsts: ", apackage.Myconsts)
}
