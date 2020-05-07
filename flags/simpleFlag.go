package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("K", true, "k flag")
	minusO := flag.Int("O", 1, "O")

	flag.Parse()

	valueK := *minusK
	valueO := *minusO

	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}
