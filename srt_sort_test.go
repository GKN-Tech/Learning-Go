package main

import (
	"reflect"
	"testing"
)

func TestSortStrings(t *testing.T) {
	input := []string{"zebra", "apple", "mango"}
	expected := []string{"apple", "mango", "zebra"}
	result := SortStrings(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
