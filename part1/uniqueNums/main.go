package main

import "fmt"

func main() {
	phoneNumbers := []int64{
		+77711111111,
		+77711111112,
		+77711111113,
		+77711111111,
		+77711111112,
	}
	uniqueNumbers := make([]int64, 0, len(phoneNumbers))
	for _, currentNum := range phoneNumbers {
		var alreadyExists = false
		for _, existingNumber := range uniqueNumbers {
			if existingNumber == currentNum {
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			uniqueNumbers = append(uniqueNumbers, currentNum)
		}
	}
	fmt.Println("Unique numbers:", uniqueNumbers)
}
