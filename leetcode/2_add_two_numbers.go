package leetcode

import (
	"github.com/ted-marozzi/algorithms/v2/datastructs"
	"golang.org/x/exp/constraints"
)

// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
// You may assume the two numbers do not contain any leading zero,
// except the number 0 itself.
// Constraints:

// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading
// zeros.
func AddTwoNumbers[T constraints.Integer](l1 *datastructs.LinkedNode[T],
	l2 *datastructs.LinkedNode[T]) *datastructs.LinkedNode[T] {

	var carry T
	rootNode := &datastructs.LinkedNode[T]{}
	currentNode := rootNode

	for l1 != nil || l2 != nil {
		// Defaults to zero if node is null
		digitOne := Deref(l1).Val
		digitTwo := Deref(l2).Val

		sum := digitOne + digitTwo + carry

		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		currentNode.Val = sum % 10
		if l1 != nil || l2 != nil {
			currentNode.Next = &datastructs.LinkedNode[T]{}
			currentNode = currentNode.Next
		}
	}

	if carry > 0 {
		currentNode.Next = &datastructs.LinkedNode[T]{Val: carry}
	}

	return rootNode
}
