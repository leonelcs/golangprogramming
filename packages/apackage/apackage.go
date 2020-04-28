package apackage

import (
	"fmt"
)

//PrintPackage method prints the string package
func PrintPackage() {
	fmt.Println("package ", privateConts)
}

const privateConts = 123

//Myconsts public constant to be available to other packages
const Myconsts = 321
