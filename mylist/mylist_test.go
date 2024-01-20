package mylist

import (
	"fmt"
	"testing"
)

func Test_myList(t *testing.T) {
	useList()
}

func Test_quickSort(t *testing.T) {
	arr := []int{2,3,1,5,6,8,11,9}
	QuickSort(arr, 0, len(arr) - 1)

	fmt.Println(arr)
}



var reverseList = function(head) {
	let current = head
	let pre = null
	while (current) {
			var t = current.next
			current.next = pre
			pre = current
			current = t
	}
	return pre
};

var reverseDoubleList = function(head) {
	let current = head
	prev = null
	while (current) {
			t = current.next
			current.next = t
			current.prev = prev
			prev = current
			current = t
	}
	return prev
}