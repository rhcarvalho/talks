package main

import "fmt"

func fibonacci() chan int {
	out := make(chan int)
	a, b := 0, 1
	go func() {
		for {
			out <- b
			a, b = b, a+b
		}
	}()
	return out
}

func main() {
	f := fibonacci()
	for x := range f {
		if x > 100 {
			break
		}
		fmt.Println(x)
	}

}

// import itertools
//
// def fibonacci():
//     a, b = 0, 1
//     while True:
//         yield b
//         a, b = b, a+b
//
// f = fibonacci()
// next(f) # 1
// [i for i in itertools.takewhile(lambda x: x < 100, f)]
// # [1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
