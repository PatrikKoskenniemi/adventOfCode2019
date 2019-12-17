package main

import (
	"fmt"
)

func main() {

	ints := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	ints[0][0] += 1
	fmt.Println(ints[0][0])
	fmt.Println(ints)
}
