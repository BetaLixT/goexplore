package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	value1, value2, err := TwoSum(nums, target)
	if err != nil {
		fmt.Println("Failed to find solutiobn")
	} else {
		fmt.Printf("Solution is %v and %v\n", value1, value2)
	}
}

type Error struct {
}

func (err Error) Error() string {
	return "No solutions found!"
}

func TwoSum(nums []int, target int) (value1 int, value2 int, err error) {
	value1 = 0
	value2 = len(nums) - 1
	for value1 != value2 {
		sum := nums[value1] + nums[value2]
		if sum > target {
			value2 -= 1
		} else if sum < target {
			value1 += 1
		} else {
			value1 = nums[value1]
			value2 = nums[value2]
			return
		}
	}
	err = Error{}
	return
}
