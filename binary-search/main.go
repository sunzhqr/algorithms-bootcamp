package main

import "fmt"

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}
	target := 8
	if result := binarySearch(arr, target); result != -1 {
		fmt.Println("The target is located in array, it's index:", result)
	} else {
		fmt.Println("The target isn't located in array")
	}
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		fmt.Printf("MIDDLE index: %d and value %d\n", middle, arr[middle])
		if arr[middle] < target {
			left = middle + 1
		} else if arr[middle] > target {
			right = middle - 1
		}
		if arr[middle] == target {
			return middle
		}
	}
	return -1
}
