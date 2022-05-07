package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		// Each outer loop iter gets the largest i-th largest
		// value in place hence - i - 1
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
