//+build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	s := []string{"Alice", "Bob", "Eve"}

	rand.Shuffle(len(s), func(i int, j int) { // HL
		s[i], s[j] = s[j], s[i] // HL
	}) // HL

	fmt.Println(s)
}
