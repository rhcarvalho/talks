package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func download(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func print_beginning(uri string) {
	result, err := download(uri)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(result[:100]))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		print_beginning("http://google.pl")
		wg.Done()
	}()
	go func() {
		print_beginning("http://google.fr")
		wg.Done()
	}()
	go func() {
		print_beginning("http://google.hu")
		wg.Done()
	}()
	wg.Wait()
}
