// 1. Create a small program which accepts two slices of colors. It must return a slice containing
// the colors that appear in either or both slices.
// c1 := []string{"Red", "Black", "White"}
// c2 := []string{"Black", "Yellow", "Orange"}
// should return in any order a slice which should have no duplicates.:
// [Red Black White Yellow Orange]

package main

import "fmt"

type methods interface {
}

func Union(slc1 []string, slc2 []string) []string {
	m := make(map[string]bool)

	outputSlice := make([]string, len(slc1)+len(slc2))
	idx := 0
	for _, val := range slc1 {
		if _, ok := m[val]; !ok {
			m[val] = true
			outputSlice[idx] = val
			idx += 1
		}
	}

	for _, val := range slc2 {
		if _, ok := m[val]; !ok {
			m[val] = true
			outputSlice[idx] = val
			idx += 1
		}
	}
	return outputSlice
}

func main() {
	c1 := []string{"Red", "Black", "White"}
	c2 := []string{"Black", "Yellow", "Orange"}

	fmt.Println(Union(c1, c2))
}
