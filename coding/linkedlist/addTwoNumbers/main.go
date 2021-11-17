package addTwoNumbers

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := false
	head := ListNode{}
	var previous *ListNode
	count := 0
	for l1 != nil && l2 != nil {
		next := ListNode{}
		var sum int
		if add {
			sum = l1.Val + l2.Val + 1
		} else {
			sum = l1.Val + l2.Val
		}

		if sum >= 10 {
			add = true
		} else {
			add = false
		}
		val := sum % 10
		if count == 0 {
			head.Val = val
			previous = &head
		} else {
			next.Val = val
			previous.Next = &next
			previous = &next
		}
		l1 = l1.Next
		l2 = l2.Next
		count++
	}
	if l1 != nil {
		for l1 != nil {
			if add {
				sum := l1.Val + 1
				l1.Val = sum % 10
				if sum >= 10 {
					add = true
				} else {
					add = false
				}
			}
			node := ListNode{}
			node.Val = l1.Val
			previous.Next = &node
			previous = &node
			l1 = l1.Next
		}
	} else if l2 != nil {
		for l2 != nil {
			if add {
				l2.Val = l2.Val + 1
				if l2.Val >= 10 {
					add = true
				} else {
					add = false
				}
			}
			node := ListNode{}
			node.Val = l2.Val % 10
			previous.Next = &node
			previous = &node
			l2 = l2.Next
		}
	}

	if add {
		tailer := ListNode{}
		tailer.Val = 1
		previous.Next = &tailer
	}
	return &head
}