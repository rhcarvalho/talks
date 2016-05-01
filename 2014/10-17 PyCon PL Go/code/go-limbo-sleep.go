package main

import "fmt"
import "time"

func main() {
	go func() {
		fmt.Println("Will you see me?")
	}()
	time.Sleep(1 * time.Second) // HL
}
