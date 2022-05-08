package leetcode

import (
	"github.com/ted-marozzi/algorithms/golang/datastructs"
	"golang.org/x/exp/constraints"
)

func MergeKLists[T constraints.Ordered](lists []*datastructs.LinkedNode[T]) *datastructs.LinkedNode[T] {
	return datastructs.MergeKLists(lists)
}
