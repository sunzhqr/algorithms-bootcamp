package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println

func maxTwo(n1, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

func main() {
	grades := []int{
		43, 47, 65, 75, 34, 31, 96, 45, 37, 23, 16,
	}
	pl("Top grades:", findTopGrades(grades, 4))
}

func findMaxUnderBoundary(grades []int, topBoundary int) int {
	currentMax := math.MinInt
	for _, grade := range grades {
		if grade < topBoundary {
			currentMax = maxTwo(currentMax, grade)
		}
	}
	return currentMax
}

func findTopGrades(inputArray []int, times int) []int {
	topGrades := make([]int, 0, times)
	previousMax := math.MaxInt
	for i := 0; i < times; i++ {
		previousMax = findMaxUnderBoundary(inputArray, previousMax)
		topGrades = append(topGrades, previousMax)
	}
	return topGrades
}
