package search

import (
	"testing"

	"github.com/ted-marozzi/algorithms/golang/sort"
)

func TestBinarySearchFoundInt(t *testing.T) {
	if !BinarySearch(1, sort.MergeSort([]int{3, 2, 1, -12, 1, 51, 12, 1})) {
		t.Error("One is in the array")
	}

	if !BinarySearch(-12, sort.MergeSort([]int{3, 2, 1, -12, 1, 51, 12, 1})) {
		t.Error("-12 is in the array")
	}

	if !BinarySearch(3, []int{3}) {
		t.Error("3 is in the array")
	}
}

func TestBinarySearchNotFoundInt(t *testing.T) {
	if BinarySearch(1, []int{}) {
		t.Error("One isn't in the empty array")
	}

	if BinarySearch(10, sort.MergeSort([]int{3, 2, 1, -12, 1, 51, 12, 1})) {
		t.Error("One is in the array")
	}

	if BinarySearch(0, sort.MergeSort([]int{3, 2, 1, -12, 1, 51, 12, 1})) {
		t.Error("0 is not in the array")
	}

	if BinarySearch(2, []int{3}) {
		t.Error("2 is not in the array")
	}
}

func TestBinarySearchFoundFloat(t *testing.T) {
	if !BinarySearch(1.0, sort.MergeSort([]float64{3, 2, 1.0, -12, 1, 51, 12, 1})) {
		t.Error("1.0 is in the array")
	}

	if !BinarySearch(-12.0, sort.MergeSort([]float64{3, 2, 1, -12.00, 1, 51, 12, 1})) {
		t.Error("-12.0 is in the array")
	}

	if !BinarySearch(3, []float64{3}) {
		t.Error("3 is in the array")
	}
}

func TestBinarySearchNotFoundFloat(t *testing.T) {
	if BinarySearch(1.0, []float64{}) {
		t.Error("One isn't in the empty array")
	}

	if BinarySearch(10.001, sort.MergeSort([]float64{0.3, 2.9181, 1.001, -112, 1.001, 5.1, 0.0001012, 1000001})) {
		t.Error("10.001 is not in the array")
	}

	if BinarySearch(0.0, sort.MergeSort([]float64{3, 2, 1, -12, 1, 51, 12, 1})) {
		t.Error("0.0 is not in the array")
	}

	if BinarySearch(3.0000002, []float64{3.0000001}) {
		t.Error("Negative 3. is in the array")
	}
}

func TestBinarySearchFoundString(t *testing.T) {
	if !BinarySearch("51", sort.MergeSort([]string{"a", "as", "asd", "-13",
		"zzzz", "51", "Hello World", "Hello World 123"})) {
		t.Error("'51' is in the array")
	}

	if !BinarySearch("Hello", sort.MergeSort([]string{"Hello", "x", "zzz"})) {
		t.Error("Hello is in the array")
	}

	if !BinarySearch("3", []string{"3"}) {
		t.Error("3 is in the array")
	}
}

func TestBinarySearchNotFoundString(t *testing.T) {
	if BinarySearch("hello", []string{"hell", "hello "}) {
		t.Error("hello isn't in the empty array")
	}

	if BinarySearch("!", sort.MergeSort([]string{"!!", "12", "!! Hello", "! !"})) {
		t.Error("! is not in the array")
	}

}
