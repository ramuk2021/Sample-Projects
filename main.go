package main

import (
	"arraysort"
	"fmt"
)

func main() {
	input := []int{6, 5, 2, 3, 9, 15, 1}
	sortHandler := arraysort.NewSortHandler()
	sortHandler.BubbleSort(input)
	fmt.Println("bubble sorted array ", input)

	input = []int{10, 34, 21, 7, -5}
	sortHandler.QuickSort(input)
	fmt.Println("quick sorted array ", input)
}
