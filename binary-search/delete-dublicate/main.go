package main

import "fmt"

func main() {
	attemptions := [7]int{
		71, 81, 64, 81, 92, 64, 81,
	}
	fmt.Println(attemptions)
	fmt.Println(attemptions[:deleteDublicate(&attemptions)])
	fmt.Println(attemptions)
}

func deleteDublicate(arr *[7]int) int {
	length := len(arr)
	i := 0
	for i < length {
		var found = false
		for j := i + 1; j < length; j++ {
			if arr[j] == arr[i] {
				found = true
				break
			}
		}
		if !found {
			i++
			continue
		}
		for k := i + 1; k < length; k++ {
			arr[k-1] = arr[k]
		}
		length--
	}
	return length
}
