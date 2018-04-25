//+build ignore

package main

import (
	"fmt"
)

type T struct {
	a int
}

func (tv T) Mv(a int) int          { return 0 } // value receiver
func (tp *T) Mp(f float32) float32 { return 1 } // pointer receiver

var t T

func main() {
	// 1 OMIT
	for _, method := range []interface{}{
		T.Mv,
		(*T).Mp,
		(*T).Mv,
	} {
		fmt.Printf("%T\n", method)
	}
	fmt.Println(T.Mv(t, 42))
}

func methodExpressionInlineComplexReceiverType() {
	// 2 OMIT
	interface{ Mv(int) int }.Mv(t, 42) // valid
	struct{ T }.Mv(struct{ T }{}, 42)  // valid, but...
}
