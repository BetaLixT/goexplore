package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("%v\n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	idx1, idx2 := 0, 0
	max := (l/2)
	if l % 2 == 0 {
	  max--
	}
	for i := 0; i < max; i++ {
		if idx1 == len(nums1) {
			idx2++
		} else if idx2 == len(nums2) {
			idx1++
		} else if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			idx2++
		}
	}
	if l%2 == 0 {
		s := 0
		if idx1 == len(nums1) {
			s = nums2[idx2]
			idx2++
		} else if idx2 == len(nums2) {
			s = nums1[idx1]
			idx1++
		} else if nums1[idx1] < nums2[idx2] {
			s = nums1[idx1]
			idx1++
		} else {
			s = nums2[idx2]
			idx2++
		}

		if idx1 == len(nums1) {
			s += nums2[idx2]
		} else if idx2 == len(nums2) {
			s += nums1[idx1]
		} else if nums1[idx1] < nums2[idx2] {
			s += nums1[idx1]
		} else {
			s += nums2[idx2]
		}
		return float64(s) / 2.0
	} else {
		s := 0
		if idx1 == len(nums1) {
			s = nums2[idx2]
		} else if idx2 == len(nums2) {
			s = nums1[idx1]
		} else if nums1[idx1] < nums2[idx2] {
			s = nums1[idx1]
		} else {
			s = nums2[idx2]
		}
		return float64(s)
	}
}
