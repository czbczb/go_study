package sorter

import (
	"fmt"
	"sort"
)

func sortCustom() {
	arr := []int{1, 4,5, 3, 2}
	sort.Ints(arr)

	// sort.Sort(sort.IntSlice(arr))
	
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}