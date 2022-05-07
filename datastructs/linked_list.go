package datastructs

import (
	"fmt"
)

type LinkedNode[T any] struct {
	Val  T
	Next *LinkedNode[T]
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
