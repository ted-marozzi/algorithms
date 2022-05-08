package leetcode

import (
	"testing"

	"github.com/ted-marozzi/algorithms/golang/datastructs"
)

func TestAddTwoNumbers(t *testing.T) {
	one := datastructs.CreateLinkedList([]int{2, 4, 3})
	two := datastructs.CreateLinkedList([]int{5, 6, 4})
	resultOne := datastructs.CreateLinkedList([]int{7, 0, 8})

	if !datastructs.CompareLinkedList(AddTwoNumbers(one, two), resultOne) {
		t.Error("1) 342 + 465 is 807")
	}

	three := datastructs.CreateLinkedList([]int{0})
	four := datastructs.CreateLinkedList([]int{0})
	resultTwo := datastructs.CreateLinkedList([]int{0})

	if !datastructs.CompareLinkedList(AddTwoNumbers(three, four), resultTwo) {
		t.Error("2) 0 + 0 is 0")
	}

	five := datastructs.CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	six := datastructs.CreateLinkedList([]int{9, 9, 9, 9})
	resultThree := datastructs.CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	if !datastructs.CompareLinkedList(AddTwoNumbers(five, six), resultThree) {
		t.Error("3) 9999999 + 9999 is 10009998")
	}
}
