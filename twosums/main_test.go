package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	first, second, err := TwoSum(nums, target)
	if err != nil {
		t.Errorf("Invalid error")
	} else if first != 2 || second != 7 {
		t.Error("Invalid response!")
	}
}
