package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Will you see me?")
	}()
}
