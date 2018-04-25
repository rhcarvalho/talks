//+build ignore

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	describe := func(v interface{}) { fmt.Fprintf(w, "%T\t%[1]v\t\n", v) }

	describe("Go 🤘")
	describe(true)
	describe(0)
	describe(0.0)
	describe('x')
	describe(0i)

	const c = 0.23 + 1.77
	var x int64 = c             // HL
	var y float32 = c           // HL
	duration := c * time.Second // HL
	describe(x)
	describe(y)
	describe(duration)
}
