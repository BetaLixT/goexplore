package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 4}
	target := 6

	index1, index2, err := TwoSum(nums, target)
	if err != nil {
		fmt.Println("Failed to find solution")
	} else {
		fmt.Printf("Solution is %v and %v\n", index1, index2)
	}
}

type Error struct {
}

func (err Error) Error() string {
	return "No solutions found!"
}

// - int with index
type IntIndex struct {
	val int
	idx int
}

func ToIntIndex (arr []int) (res []IntIndex) {
	res = make([]IntIndex, len(arr))
	for i, val := range arr {
		res[i] = IntIndex{val: val, idx: i}
	}
	return
}


func TwoSum(nums []int, target int) (index1 int, index2 int, err error) {
	
	numsi := ToIntIndex(nums)
	sort.Slice(numsi, func (i, j int) bool {
		return numsi[i].val < numsi[j].val
	})
	
	index1 = 0
	index2 = len(numsi) - 1
	for index1 != index2 {
		sum := numsi[index1].val + numsi[index2].val
		if sum > target {
			index2 -= 1
		} else if sum < target {
			index1 += 1
		} else {
			index1 = numsi[index1].idx
			index2 = numsi[index2].idx
			return
		}
	}
	err = Error{}
	return
}
