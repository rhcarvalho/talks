package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	urls := []string{
		"https://google.fr",
		"https://google.hu",
		"https://google.pl",
	}
	var code int32
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		url := url
		go func() {
			if err := download(url); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				// set exit code to 1 in case of error
				atomic.CompareAndSwapInt32(&code, 0, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	os.Exit(int(code))
}

func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%v: %v (%d bytes)\n", url, resp.Status, len(b))
	return nil
}
