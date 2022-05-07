package datastructs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type LinkedNode[T any] struct {
	Val  T
	Next *LinkedNode[T]
}

func MergeKLists[T constraints.Ordered](lists []*LinkedNode[T]) *LinkedNode[T] {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return MergeLists(lists[0], lists[1])
	}

	middle := len(lists) / 2

	return MergeLists(MergeKLists(lists[:middle]), MergeKLists(lists[middle:]))
}

// Assumes Sorted
func MergeLists[T constraints.Ordered](rootOne, rootTwo *LinkedNode[T]) *LinkedNode[T] {
	result := &LinkedNode[T]{}
	currentNode := result

	if rootOne == nil && rootTwo == nil {
		return nil
	}

	for rootOne != nil || rootTwo != nil {
		// If we finish one of the slices append whatever is left
		if rootOne == nil {
			currentNode.Next = rootTwo.Next
			currentNode.Val = rootTwo.Val
			return result
		}

		if rootTwo == nil {
			currentNode.Next = rootOne.Next
			currentNode.Val = rootOne.Val
			return result
		}

		// Append the smaller of the numbers
		if rootOne.Val <= rootTwo.Val {
			currentNode.Val = rootOne.Val
			rootOne = rootOne.Next
		} else if rootOne.Val > rootTwo.Val {
			currentNode.Val = rootTwo.Val
			rootTwo = rootTwo.Next
		}
		currentNode.Next = &LinkedNode[T]{}
		currentNode = currentNode.Next
	}
	return result
}

func CreateLinkedList[T any](slice []T) *LinkedNode[T] {
	var val T
	root := &LinkedNode[T]{Val: val}
	currentNode := root
	for i, val := range slice {
		currentNode.Val = val
		if i != len(slice)-1 {
			currentNode.Next = &LinkedNode[T]{}
			currentNode = currentNode.Next
		}
	}
	return root
}

func PrintLinkedList[T any](root *LinkedNode[T]) {
	fmt.Println("List of nodes: ")
	for root != nil {
		fmt.Print(root.Val)
		root = root.Next
	}
	fmt.Println("")
}

func CompareLinkedList[T comparable](
	rootOne *LinkedNode[T],
	rootTwo *LinkedNode[T],
) bool {
	for rootOne != nil && rootTwo != nil {
		if rootOne.Val != rootTwo.Val {
			return false
		}
		rootOne = rootOne.Next
		rootTwo = rootTwo.Next
	}

	// If all the values were the same and we exhausted both lists
	// then return true, otherwise false
	return rootOne == nil && rootTwo == nil
}
