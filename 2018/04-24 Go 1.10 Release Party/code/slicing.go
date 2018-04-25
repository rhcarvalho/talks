//+build ignore

package main

import "fmt"

func main() {
	s := []byte("abracadabra")
	s1 := s[0:4] // HL
	s1 = append(s1, 'Z')
	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s) // abracadabra?

	s = []byte("abracadabra")
	s2 := s[0:4:4] // HL
	s2 = append(s2, 'Z')
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", s) // abracadabra?
}
