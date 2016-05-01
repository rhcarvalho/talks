package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// typedefs OMIT

type GitHubSearchResponse struct {
	Items []*Repo
}

type Repo struct {
	Fullname string `json:"full_name"`
	Stars    int    `json:"stargazers_count"`
}

// -typedefs OMIT

func (sr *GitHubSearchResponse) print() {
	fmt.Printf(" #  %-20s %s\n", "Repo URL", "Stars")
	fmt.Println(strings.Repeat("-", 30))
	for i, r := range sr.Items {
		fmt.Printf("%02d. %-20s %d\n", i+1, r.Fullname, r.Stars)
	}
}

func main() {
	// Search GitHub repositories
	// req OMIT
	// massign OMIT
	query, language := "python", "go"
	// -massign OMIT
	url := `https://api.github.com/search/repositories?q=` + query +
		`+language:` + language + `&sort=stars&order=desc&per_page=3`
	resp, err := http.Get(url) // HLreq
	checkErr(err)
	defer resp.Body.Close() // HLdefer
	// -req OMIT

	// Decode JSON response
	// json OMIT
	dec := json.NewDecoder(resp.Body)
	var sr GitHubSearchResponse
	err = dec.Decode(&sr)
	checkErr(err)
	// -json OMIT

	// Print results
	sr.print()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
		// fmt.Print(...)
		// os.Exit(1)
	}
}
