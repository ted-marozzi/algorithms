package sort

import "golang.org/x/exp/constraints"

func SimpleSort[T constraints.Ordered](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
