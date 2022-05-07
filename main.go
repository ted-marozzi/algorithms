package main

import (
	"fmt"

	"github.com/ted-marozzi/algorithms/v2/search"
	"github.com/ted-marozzi/algorithms/v2/sort"
)

func main() {
	// ints
	intSlice := []int{1, -10, 1, 100, 12, 4, 52, -10000200, 10012, 1, 0}
	sortedIntSlice := sort.MergeSort(intSlice)

	// strings
	stringSlice := []string{
		"b", "a", "1233", "-123", "Hello World", "z", "!!", "1q2",
	}
	sortedStringSlice := sort.MergeSort(stringSlice)

	// Sorts
	sort.PrintSliceSorted(intSlice, sortedIntSlice)
	sort.PrintSliceSorted(stringSlice, sortedStringSlice)

	// Search
	target := "Hello World"
	if search.BinarySearch(target, sortedStringSlice) {
		fmt.Printf("%s is found in %+v\n", target, sortedStringSlice)
	}

	targetTwo := "HelloWorld"
	if !search.BinarySearch(targetTwo, sortedStringSlice) {
		fmt.Printf("%s is NOT found in %+v\n", targetTwo, sortedStringSlice)
	}

}
