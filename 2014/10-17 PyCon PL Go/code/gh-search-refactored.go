package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	GitHubSearchResponse struct {
		Items []*Repo
	}
	Repo struct {
		Fullname string `json:"full_name"`
		Stars    int    `json:"stargazers_count"`
	}
)

func main() {
	sr, err := searchPopularByLanguage("python", "go") // HLrefactor
	checkErr(err)
	sr.print(os.Stdout) // HLrefactor
}

// search OMIT

func searchPopularByLanguage(query, language string) (*GitHubSearchResponse, error) {
	queryString := `q=` + query + `+language:` + language + `&sort=stars&order=desc&per_page=3`
	return searchRepo(queryString)
}

func searchRepo(queryString string) (*GitHubSearchResponse, error) {
	var sr *GitHubSearchResponse
	url := `https://api.github.com/search/repositories?` + queryString
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

// -search OMIT

func (sr *GitHubSearchResponse) print(w io.Writer) {
	width := 30
	fmt.Fprintf(w, " #  %-*s Stars\n", width, "Repo URL")
	fmt.Fprintln(w, strings.Repeat("-", width+10))
	for i, r := range sr.Items {
		fmt.Fprintf(w, "%02d. %-*s %5d\n",
			i+1, width, r.Fullname, r.Stars)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
