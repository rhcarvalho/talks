package main

import "fmt"

func main() {
	loopRange()
}

func loopRange() {
	for i, el := range []int{11, 7, 87} {
		// like Python:
		// for i, el in enumerate([11, 7, 87]):
		fmt.Println(i, el)
	}
}

func loopC() {
	for i := 0; i < 100; i++ {
		// like C
	}
}

const condition = true

func loopWhile() {
	for condition {
		// while loop
	}
}

func loopInf() {
	for {
		// infinite loop
	}
}
