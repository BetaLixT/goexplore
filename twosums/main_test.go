package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{7, 11, 2, 15}
	target := 9

	first, second, err := TwoSum(nums, target)
	if err != nil {
		t.Errorf("Invalid error")
	} else if (first != 0 || second != 2 ) &&  (first != 2 || second != 0 ){
		t.Error("Invalid response!")
	}
}
