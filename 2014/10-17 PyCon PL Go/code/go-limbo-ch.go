package main

import "fmt"

func main() {
	done := make(chan bool) // HL
	go func() {
		fmt.Println("Will you see me?")
		done <- true // HL
	}()
	<-done // HL
}
