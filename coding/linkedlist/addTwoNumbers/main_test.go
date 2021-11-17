package addTwoNumbers

import "testing"

type input struct {
	a1 *ListNode
	a2 *ListNode
	expected *ListNode
}

func TestAddTwoNumbers(t *testing.T)  {
	at1 := ListNode{Val: 2}
	at2 := ListNode{Val: 4}
	at1.Next = & at2
	at3 := ListNode{Val: 3}
	at2.Next = &at3

	bt1 := ListNode{Val: 5}
	bt2 := ListNode{Val: 6}
	bt1.Next = &bt2
	bt3 := ListNode{Val: 4}
	bt2.Next = &bt3

	et1 := ListNode{Val: 7}
	et2 := ListNode{Val: 0}
	et1.Next = &et2
	et3 := ListNode{Val: 8}
	et2.Next = &et3
	inputs := []input{
		{&at1, &bt1, &et1},
	}
	for _, i := range inputs {
		result := addTwoNumbers(i.a1, i.a2)
		temp := i.expected
		for temp != nil {
			if result == nil || result.Val != temp.Val {
				t.Fatalf("Failed result, expected: %v", temp.Val)
			}
			temp = temp.Next
			result = result.Next
		}
	}
}
