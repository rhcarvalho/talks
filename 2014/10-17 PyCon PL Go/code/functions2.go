package main

import "fmt"

func max(args ...int) int { // HL
	m := 0
	for _, x := range args { // HL
		if x > m {
			m = x
		}
	}
	return m
}

func main() {
	fmt.Println(max(6, 9))
	fmt.Println(max(8, 1))
	fmt.Println(max(8, 1, 45))
	fmt.Println(max(7, 13, 2, 5, 42))
}
