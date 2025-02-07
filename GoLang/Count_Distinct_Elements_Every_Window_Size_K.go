package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 3, 4, 2, 3}
	k := 4
	myMap := make(map[int]int)
	for i := 0; i < k; i++ {
		myMap[arr[i]]++
	}
	fmt.Print(len(myMap), " ")
	for i := k; i < len(arr); i++ {
		lastKey := arr[i-k]
		v, _ := myMap[lastKey]
		if v == 1 {
			delete(myMap, lastKey)
		} else {
			myMap[lastKey]--
		}
		myMap[arr[i]]++
		fmt.Print(len(myMap), " ")
	}
}
