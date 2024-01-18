package sorter

import (
	"fmt"
	"testing"
)


func Test_bubbleSort(t *testing.T) {
	slice2 := []int{1, 7, 3, 4, 0, 2, 10, 5}
	BubbleSort(slice2)

	fmt.Println(slice2)
}