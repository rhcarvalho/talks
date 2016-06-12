package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("bash", "-c", "while true; do date && sleep 1; done")
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	timeout := 5 * time.Second
	time.AfterFunc(timeout, func() { // HL
		cmd.Process.Kill() // HL
	}) // HL

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
