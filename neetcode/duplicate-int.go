package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 3}

	hashMap := make(map[int]bool)

	for _, num := range array {
		_, ok := hashMap[num]

		if ok {
			fmt.Println(true)
			return
		} else {
			hashMap[num] = true
		}
	}
	fmt.Println(false)
}
