package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

// List2IntArray convert List to int array
func List2IntArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
