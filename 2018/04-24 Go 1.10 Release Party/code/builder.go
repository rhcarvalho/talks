//+build ignore

package main

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	_ bytes.Buffer
	_ strings.Builder
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("Go Party!")
	fmt.Println(b.String())

}
