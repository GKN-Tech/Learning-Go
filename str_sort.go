package main

import "sort"

// SortStrings sorts a slice of strings in ascending order.
func SortStrings(input []string) []string {
	sort.Strings(input)
	return input
}
