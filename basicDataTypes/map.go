package main

import "fmt"

func main() {

	iMap := make(map[rune]int)

	sentence := "O rato roeu a roupa do rei de roma"

	for _, char := range sentence {
		_, ok := iMap[char]
		if ok {
			iMap[char]++
		} else {
			iMap[char] = 1
		}
	}
	for key, value := range iMap {
		fmt.Println("key: ", string(key), " frequency: ", value)
	}

	//it will panic
	aMap := map[string]int{}
	// var aMap map[string]int
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1

}
