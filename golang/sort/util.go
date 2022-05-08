package sort

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func PrintSliceSorted[T constraints.Ordered](unsorted, sorted []T) {
	fmt.Printf(
		"Original slice: %+v\nSorted slice:   %+v\n",
		unsorted,
		sorted,
	)
}

func CheckSorted[T constraints.Ordered](slice []T) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
