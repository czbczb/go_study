package useslice

import (
	"encoding/json"
	"fmt"
	"time"
)

func DemoSlice() {
// 如果切片的容量小于 1024，则将容量扩大为原来的 2 倍。
// 如果切片的容量大于等于 1024，则将容量扩大为原来的 1.25 倍。


	slice1 := make([]int , 3, 10)
	slice2 := make([]int , 3, 3)

	// 当slice2的capacity能够承载所有元素，则修改slice2，否则生成新的数组
	slice22 := append(slice2, 2,3,4,5)
	slice3 := []int{1,2,3,4}
	slice4 := [...]int{1,2,3,4}


	// for i,v := range slice1 {
	// 	fmt.Println(i, v)
	// }

	for i :=0; i < len(slice1); i++ {
		fmt.Println(i, slice1[i])
	}
	time.Sleep(1*time.Second)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2, slice22, "cap:", cap(slice2), cap(slice22))
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)


	slicetest1 := []int{1,2,3,4}
	slicetest2 := []int{5,6}

	// copy(slicetest1, slicetest2)
	copy(slicetest2, slicetest1)

	fmt.Println(slicetest2)


}

func deleteElement(array []int, index int) []int {
	if index < 0 || index >= len(array) {
			return array
	}

	// 截取前半部分
	front := array[:index]

	// 截取后半部分
	back := array[index+1:]

	// 重新组合
	return append(front, back...)
}


func DemoArray() {

	type MyMessage struct {
		Name string
		Age int
	}
	array1 := [5]int{1,2,3,4, 5}
	array2 := [5]MyMessage{}

	fmt.Println(array1)


	jsonStr, err := json.Marshal(array2)
	if err != nil {
		fmt.Println("jsonStr:", err)
	}
	fmt.Println(string(jsonStr))


	// 不定参数的函数调用
	DemoFunc(array1[0:]...)
}


func DemoFunc(args ...int) {

}

