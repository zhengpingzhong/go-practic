package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)
	// 判断是否已排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)
}

//strings: [a b c]
//ints:  [2 4 7]
//sorted:  true
