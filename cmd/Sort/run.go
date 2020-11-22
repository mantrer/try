package sort

import "fmt"

func searchMin(a []int) (int, int) {
	var ind int
	var v int

	min := a[0]

	for ind, v = range a {
		if v < min {
			min = v
		}
	}
	return ind - 1, min
}

func sortArr(a []int) []int {
	var arr []int
	count := len(a)
	for i := 0; i < count-1; i++ {
		ind, val := searchMin(a)
		fmt.Println(ind, val)
		fmt.Println("?", val)
		arr = append(arr, val)
		// delete min element from slice
		anew := append(a[:ind], a[ind+1])
		a = anew
		fmt.Println(">>>", a)
	}
	fmt.Println(arr)
	return arr
}
