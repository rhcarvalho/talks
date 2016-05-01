package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(6, 9))
	fmt.Println(max(8, 1))
}
