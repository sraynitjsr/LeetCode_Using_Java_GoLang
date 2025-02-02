package main

import "fmt"

var myMap map[string]struct{}

func fillMap() {
	myMap = make(map[string]struct{})
	myMap["A"] = struct{}{}
	myMap["a"] = struct{}{}
	myMap["E"] = struct{}{}
	myMap["e"] = struct{}{}
	myMap["I"] = struct{}{}
	myMap["i"] = struct{}{}
	myMap["O"] = struct{}{}
	myMap["o"] = struct{}{}
	myMap["U"] = struct{}{}
	myMap["u"] = struct{}{}
}

func main() {
	myStr := "abciiidef"
	K := 3
	fillMap()
	finalCount := 0
	windowCount := 0
	for i := 0; i < K; i++ {
		_, present := myMap[string(myStr[i])]
		if present {
			windowCount++
		}
	}
	finalCount = windowCount
	for i := K; i < len(myStr); i++ {
		_, presentAtI := myMap[string(myStr[i])]
		if presentAtI {
			windowCount++
		}
		_, presentAtIMinusK := myMap[string(myStr[i-K])]
		if presentAtIMinusK {
			windowCount--
		}
		if windowCount > finalCount {
			finalCount = windowCount
		}
	}
	fmt.Println(finalCount)
}
