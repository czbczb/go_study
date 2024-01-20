package mylist

import (
	"container/list"
	"fmt"
)

func useList () {
	list1 := list.New()

	list1.PushFront("a") // [a]
	list1.PushBack("b")  // [a, b]
	list1.InsertBefore("c", list1.Back()) // [a, c, b]
	list1.InsertAfter("d", list1.Front()) // [a,d,c,b]

	// 取中间元素
	getListValueN(list1, 2) // d


	// 移动
	// list1.MoveAfter(list1.Front(), list1.Back())
	// list1.MoveBefore(list1.Front(), list1.Back())
	// list1.MoveToBack(list1.Front()) // 把某个元素移动到最后面
	// list1.MoveToFront(list1.Back()) // 把某个元素移动到最前面

	// 遍历		
	for e:= list1.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
	}
}



func getListValueN(list1 *list.List, n int) {
	if n > list1.Len() || n < 0 {
		fmt.Println("n 的值不存在链表中")
		return
	}

	var current *list.Element
	if n == 1 {
		current = list1.Front()
	}else if n == list1.Len() {
		current = list1.Back()
	}else {
		current := list1.Front()
		for i := 1; i < n; i++ {
			current = current.Next()
		}
	}

	fmt.Println("current value:", current.Value)
}


func reverseList(list1 *list.List) {
	// header := list1.Front()
	// end := list1.Back()

}

func QuickSort(arr []int, left, right int) {
	if(left > right) {
		return
	}
	base := arr[left]
	i := left
	j := right
	
	for i< j && arr[j] >= base {
		j--
	}
	for i< j && arr[i] <= base {
		i++
	}
	// 找到了
	if i < j {
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
	}

	arr[left] = arr[i]
	arr[i] = base

	QuickSort(arr, left, i-1)
	QuickSort(arr, i + 1, right)
}