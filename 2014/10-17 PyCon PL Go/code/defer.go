package main

import "fmt"

func main() {
	fmt.Println("A")
	defer func() {
		fmt.Println("Deferred 1")
	}()
	fmt.Println("B")
	defer func() {
		fmt.Println("Deferred 2")
	}()
}
