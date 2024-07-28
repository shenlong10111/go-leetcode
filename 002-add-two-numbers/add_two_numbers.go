package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	currentResult := result
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currentResult.Next = &ListNode{Val: sum % 10}
		currentResult = currentResult.Next
		carry = sum / 10
	}

	return result.Next
}

func main() {
	l1 := &ListNode{Val: 9}
	l2 := &ListNode{Val: 9}
	fmt.Println(addTwoNumbers(l1, l2))

	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 4}
	fmt.Println(addTwoNumbers(l3, l4))
}
