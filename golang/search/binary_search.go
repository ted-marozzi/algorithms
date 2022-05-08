package search

import "golang.org/x/exp/constraints"

// Returns true if the target exists in the slice assuming the slice is sorted
func BinarySearch[T constraints.Ordered](target T, sliceToSearch []T) bool {

	// Not found case
	if len(sliceToSearch) == 0 {
		return false
	}

	middle := len(sliceToSearch) / 2

	if sliceToSearch[middle] < target {
		return BinarySearch(target, sliceToSearch[middle+1:])
	} else if sliceToSearch[middle] > target {
		return BinarySearch(target, sliceToSearch[:middle])
	}
	// If the current index isn't > or < the target it must == target, hence
	// we found it
	return true
}
