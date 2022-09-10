package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	res := addTwoNumbersPipeline(l1, l2)
	for res != nil {
		fmt.Printf("%d\n", res.Val)
		res = res.Next
	}
}

type Vals struct {
	v1 int
	v2 int
}

func addTwoNumbersPipeline(l1 *ListNode, l2 *ListNode) *ListNode {
	carrychan := make(chan int, 1000)
	go summer(l1, l2, carrychan)
	return resser(carrychan)
}

func summer(l1 *ListNode, l2 *ListNode, carrychan chan int) {
	for l1 != nil || l2 != nil {
		s := 0
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		carrychan <- s
	}
	close(carrychan)
}

func resser(carrychan chan int) *ListNode {
	carry := 0
	sum := 0

	// - first iter
	itr := &ListNode{}
	res := itr
	val, ok := <-carrychan
	sum = val + carry
	itr.Val = sum % 10
	carry = sum / 10

	for {
		val, ok = <-carrychan
		if !ok {
			break
		}
		itr.Next = &ListNode{}
		itr = itr.Next
		sum = val + carry
		itr.Val = sum % 10
		carry = sum / 10
	}
	if carry != 0 {
		itr.Next = &ListNode{
			Val: 1,
		}
	}
	return res
}

func addTwoNumbersBasic(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	v1 := 0
	v2 := 0
	sum := 0

	// - first iter
	itr := &ListNode{}
	res := itr

	if l1 != nil {
		v1 = l1.Val
		l1 = l1.Next
	} else {
		v1 = 0
	}
	if l2 != nil {
		v2 = l2.Val
		l2 = l2.Next
	} else {
		v2 = 0
	}
	sum = v1 + v2 + carry
	itr.Val = sum % 10
	carry = sum / 10
	// - other iters
	for l1 != nil || l2 != nil {
		itr.Next = &ListNode{}
		itr = itr.Next

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		fmt.Printf("%d + %d + %d\n", v1, v2, carry)
		sum = v1 + v2 + carry
		itr.Val = sum % 10
		carry = sum / 10
	}
	if carry != 0 {
		itr.Next = &ListNode{
			Val: 1,
		}
	}
	return res
}
