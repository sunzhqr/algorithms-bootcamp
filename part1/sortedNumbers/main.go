package main

import "fmt"

func main() {
	sortedNumbers := []int{
		1000, 1000, 1000,
		1001,
		1002, 1002, 1002,
		1003, 1003,
		1004,
		1005, 1005,
	}
	uniqueNumbers := make([]int, 0, len(sortedNumbers))
	prevNum := sortedNumbers[0]

	for i := 1; i < len(sortedNumbers); i++ {
		if prevNum != sortedNumbers[i] {
			uniqueNumbers = append(uniqueNumbers, prevNum)
			prevNum = sortedNumbers[i]
		}
	}
	uniqueNumbers = append(uniqueNumbers, prevNum)

	fmt.Println("Unique numbers:", uniqueNumbers)
}
