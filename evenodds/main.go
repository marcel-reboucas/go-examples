package main

import (
	"fmt"
)

func main() {
	n := makeSlice(10)

	for _, v := range n {
		if v%2 == 0 {
			fmt.Printf("%v is even \n", v)
		} else {
			fmt.Printf("%v is odd \n", v)
		}
	}
}

func makeSlice(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}
	return slice
}
