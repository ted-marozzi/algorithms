package sort

import (
	"golang.org/x/exp/constraints"
)

// Recursively cut the slice in half until we are length 0, 1 or 2
// at that point it is trivial to sort so we sort and return.
// These sorted slices get passed to Merge which will merge them into one
// sorted slice. This happens recursively until all the slices are merged into
// the final slice.
func MergeSort[T constraints.Ordered](slice []T) []T {

	// Odd case
	if len(slice) <= 1 {
		return slice
	}

	// Even case
	if len(slice) == 2 {
		if slice[0] <= slice[1] {
			return slice
		}
		return []T{slice[1], slice[0]}
	}

	middle := len(slice) / 2
	return Merge(
		MergeSort(slice[:middle]),
		MergeSort(slice[middle:]),
	)
}

// Merge two sorted arrays into one sorted array
func Merge[T constraints.Ordered](sliceOne, sliceTwo []T) []T {

	result := make([]T, 0, len(sliceOne)+len(sliceTwo))

	i, j := 0, 0
	for k := 0; k < len(sliceOne)+len(sliceTwo); k++ {
		// If we finish one of the slices append whatever is left
		if i >= len(sliceOne) || j >= len(sliceTwo) {
			result = append(result, sliceOne[i:]...)
			result = append(result, sliceTwo[j:]...)
			return result
		}

		// Append the smaller of the numbers
		if sliceOne[i] <= sliceTwo[j] {
			result = append(result, sliceOne[i])
			i++
		} else if sliceOne[i] > sliceTwo[j] {
			result = append(result, sliceTwo[j])
			j++
		}

	}

	return result
}
