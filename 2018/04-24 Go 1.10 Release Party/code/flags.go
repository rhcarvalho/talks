//+build ignore

package main

import (
	"flag"
	"fmt"
)

func main() {
	// BEGIN OMIT
	s := flag.String("s", "", "this is a description\nthat spans multiple lines")
	x := flag.Int("x", 0, "short description")
	// END OMIT
	flag.Parse()
	fmt.Printf("s = %q\nx = %d\n", *s, *x)
}
