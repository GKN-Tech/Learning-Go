package main

import "sort"

// SortIntegers sorts a slice of integers in ascending order.
func SortIntegers(input []int) []int {
	sort.Ints(input)
	return input
}
