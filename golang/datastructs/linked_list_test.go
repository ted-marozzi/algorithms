package datastructs

import "testing"

func TestMergeLists(t *testing.T) {
	result := MergeLists(
		CreateLinkedList([]int{-10, -1, 2, 3}),
		CreateLinkedList([]int{-100, 2, 15}),
	)

	if !CompareLinkedList(
		result, CreateLinkedList([]int{-100, -10, -1, 2, 2, 3, 15})) {
		t.Error("The lists were merged incorrectly")
	}
}

func TestMergeKLists(t *testing.T) {
	result := MergeKLists([]*LinkedNode[int]{
		CreateLinkedList([]int{-10, -1, 2, 3}),
		CreateLinkedList([]int{-100, 2, 15}),
		CreateLinkedList([]int{-12, 3, 15}),
		CreateLinkedList([]int{0}),
	})

	if !CompareLinkedList(
		result, CreateLinkedList([]int{-100, -12, -10, -1, 0, 2, 2, 3, 3, 15, 15})) {
		t.Error("The lists were merged incorrectly")
	}
}
