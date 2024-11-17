package main

import (
	"reflect"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}
	result := SortIntegers(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}