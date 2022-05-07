package sort

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeSortInt(t *testing.T) {
	inputOne := []int{-1, 0, 1, 2, 3}
	resultOne := MergeSort(inputOne)

	if !cmp.Equal(inputOne, resultOne) {
		t.Error("Sorting a sorted list should not change it")
	}

	if !CheckSorted(MergeSort([]int{
		50, 12, -1, 2, 4, 5, 6, 1, -10000, 12222, 1, 0, 12, 55, 11, 1, 1, 11,
		14, -123415})) {
		t.Error("Should be sorted two")
	}

	if !CheckSorted(MergeSort([]int{})) {
		t.Error("Should be sorted three")
	}

	if !CheckSorted(MergeSort([]int{0})) {
		t.Error("Should be sorted four")
	}

	if !CheckSorted(MergeSort([]int{-1, 1})) {
		t.Error("Should be sorted five")
	}

	if !CheckSorted(MergeSort([]int{0, 2})) {
		t.Error("Should be sorted six")
	}
}

func TestMergeSortFloat(t *testing.T) {
	if !CheckSorted(
		MergeSort(
			[]float64{
				50.0, 12.12, -1.0, 2.1, 4.1,
				5.1323, 6.1221, 0.1, -1000.0,
				1222.2, 1.00000012002, 0.000000001, 12.0, -34.3, 11.34, 1.3, 1.6, 11.0,
				1.4, -12.3415,
			})) {
		t.Error("Should be sorted float")
	}
}

func TestMergeString(t *testing.T) {

	if !CheckSorted(
		MergeSort(
			[]string{
				"b", "a", "1233", "-123", "helloWorld", "z", "!!", "1q2",
			})) {
		t.Error("Should be sorted float")
	}
}
