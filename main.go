package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

	// myString := ""
	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	myString = "Please give me one argument!"
	// } else {
	// 	myString = arguments[1]
	// }

	// io.WriteString(os.Stdout, "This is Standard output\n")
	// io.WriteString(os.Stderr, myString)
	// io.WriteString(os.Stderr, "\n")
}
