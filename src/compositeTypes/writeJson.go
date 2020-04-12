package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJson(filename *os.File, key interface{}) {
	encodeJson := json.NewEncoder(filename)
	err := encodeJson.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Leonel",
		Surname: "Candido",
		Tel: []Telephone{
			Telephone{
				Mobile: true,
				Number: "1234-5678"},
			Telephone{
				Mobile: true,
				Number: "1234-acbd"},
			Telephone{
				Mobile: false,
				Number: "abcc-567"}}}

	saveToJson(os.Stdout, myRecord)
}
