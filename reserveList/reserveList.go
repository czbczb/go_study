package reserveList

import (
	"container/list"
)

func ReserveList(list1 *list.List, left, right) {
	dummyNode := &list.Element{Value: -1}
	dummyNode.Next = list1
	prev := dummyNode

	for i := 0; i < left - 1; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for i := 0; i < right - left; i++ {
		t := cur.Next

		cur.Next = t.Next.Next

		t.Next = prev.Next

		prev.Next = t
	}

	return dummyNode.Next
}


// 爬楼梯
func ClimbStairs(n int) int {
	if n <= 2 {
			return n
	}
	
	first, second := 1, 2
	for i := 3; i <= n; i++ {
			first, second = second, first+second
	}
	
	return second
}

