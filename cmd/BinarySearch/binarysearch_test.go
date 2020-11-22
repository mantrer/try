package binarysearch

import (
	"fmt"
	"testing"
)

func TestSortStr(t *testing.T) {
	fmt.Println("===")
	SortStr([]string {"sss", "ggggg", "wwww", "llllll"})
	fmt.Println("===")
}

func TestBin(t *testing.T) {
	fmt.Println(checkBin([]int{1, 2, 3, 4, 5}, 1))  // true
	fmt.Println(checkBin([]int{1, 2, 3, 4, 5}, -1)) //false
}