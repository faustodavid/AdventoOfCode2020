package main

import (
	"testing"
)

func TestGetMultiplicationOf2EntriesThatSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expectedResult := uint64(514579)

	actualResult := GetMultiplicationOf2EntriesThatSum2020(entries)

	if expectedResult != actualResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actualResult, expectedResult)
	}
}

func TestGetMultiplicationOf3EntriesThatSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expectedResult := uint64(241861950)

	actualResult := GetMultiplicationOf3EntriesThatSum2020(entries)

	if expectedResult != actualResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actualResult, expectedResult)
	}
}
