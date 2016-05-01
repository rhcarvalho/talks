package main

import "fmt"

func main() {
	done := make(chan struct{}) // HL
	go func() {
		fmt.Println("Will you see me?")
		done <- struct{}{} // HL
	}()
	<-done
}
