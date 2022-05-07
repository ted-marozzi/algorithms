package main

import (
	"github.com/ted-marozzi/algorithms/v2/datastructs"
	"github.com/ted-marozzi/algorithms/v2/leetcode"
	_ "github.com/ted-marozzi/algorithms/v2/search"
	_ "github.com/ted-marozzi/algorithms/v2/sort"
)

func main() {

	one := datastructs.CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	datastructs.PrintLinkedList(one)

	two := datastructs.CreateLinkedList([]int{9, 9, 9, 9})
	datastructs.PrintLinkedList(two)

	result := leetcode.AddTwoNumbers(one, two)
	datastructs.PrintLinkedList(result)

}
