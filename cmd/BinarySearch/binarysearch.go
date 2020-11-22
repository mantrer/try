package binarysearch

import (
	"fmt"
	"sort"
)

// SortStr tst
func SortStr(s []string) {
	sort.Strings(s)
	fmt.Println(s)
}

func checkBin(list []int, i int) bool {
	indMin := 0
	indMax := len(list)
	for indMax != 0 {
		indProb := (indMax - indMin) / 2
		prob := list[indProb]

		if prob == i {
			return true
		}

		if i > prob {
			indMin = indProb
		} else {
			indMax = indProb
		}
	}
		return false
}
