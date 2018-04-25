//+build ignore

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime version:", runtime.Version())

	a := []string{"Shifting", "untyped", "constant", "1.0", "converts", "to", "int"}
	var s uint = 2

	fmt.Println(a[1.0<<s]) // 1.0 has type int, 1.0<<s is int == 4 // HL

	// Output (go1.9 and earlier):
	// invalid operation: 1 << s (shift of type float64)

	// Output (go1.10):
	// converts
}
