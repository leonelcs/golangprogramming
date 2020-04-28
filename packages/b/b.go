package b

import (
	"fmt"

	"github.com/leonelcs/golangprogramming/packages/a"
)

func init() {
	fmt.Println("initiing b")
}

//PrintB is a method that do some prints
func PrintB() {
	fmt.Println("Printing b")
	a.PrintA()
}
