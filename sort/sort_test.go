package sort

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Simple
func TestSimpleSortInt(t *testing.T) {
	SortTestHelperInt(t, SimpleSort[int])
}

func TestSimpleSortFloat(t *testing.T) {
	SortTestHelperFloat(t, SimpleSort[float64])
}

func TestSimpleSortString(t *testing.T) {
	SortTestHelperString(t, SimpleSort[string])
}

// Merge
func TestMergeSortInt(t *testing.T) {
	SortTestHelperInt(t, MergeSort[int])
}

func TestMergeSortFloat(t *testing.T) {
	SortTestHelperFloat(t, MergeSort[float64])
}

func TestMergeSortString(t *testing.T) {
	SortTestHelperString(t, MergeSort[string])
}

// Bubble
func TestBubbleSortInt(t *testing.T) {
	SortTestHelperInt(t, BubbleSort[int])
}

func TestBubbleSortFloat(t *testing.T) {
	SortTestHelperFloat(t, BubbleSort[float64])
}

func TestBubbleSortString(t *testing.T) {
	SortTestHelperString(t, BubbleSort[string])
}

func SortTestHelperInt(t *testing.T, sortingFunc func([]int) []int) {
	inputOne := []int{-1, 0, 1, 2, 3}
	resultOne := sortingFunc(inputOne)

	if !cmp.Equal(inputOne, resultOne) {
		t.Error("Sorting a sorted list should not change it")
	}

	if !CheckSorted(sortingFunc([]int{
		50, 12, -1, 2, 4, 5, 6, 1, -10000, 12222, 1, 0, 12, 55, 11, 1, 1, 11,
		14, -123415})) {
		t.Error("Should be sorted two")
	}

	if !CheckSorted(sortingFunc([]int{})) {
		t.Error("Should be sorted three")
	}

	if !CheckSorted(sortingFunc([]int{0})) {
		t.Error("Should be sorted four")
	}

	if !CheckSorted(sortingFunc([]int{-1, 1})) {
		t.Error("Should be sorted five")
	}

	if !CheckSorted(sortingFunc([]int{0, 2})) {
		t.Error("Should be sorted six")
	}
}

func SortTestHelperFloat(t *testing.T, sortingFunc func([]float64) []float64) {
	if !CheckSorted(
		sortingFunc(
			[]float64{
				50.0, 12.12, -1.0, 2.1, 4.1,
				5.1323, 6.1221, 0.1, -1000.0,
				1222.2, 1.00000012002, 0.000000001, 12.0, -34.3, 11.34, 1.3, 1.6, 11.0,
				1.4, -12.3415,
			})) {
		t.Error("Should be sorted floats")
	}
}

func SortTestHelperString(t *testing.T, sortingFunc func([]string) []string) {
	if !CheckSorted(
		sortingFunc(
			[]string{
				"b", "a", "1233", "-123", "helloWorld", "z", "!!", "1q2",
			})) {
		t.Error("Should be sorted strings")
	}
}
