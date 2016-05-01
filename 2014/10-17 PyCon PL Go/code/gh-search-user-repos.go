package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var requestCounter int

type (
	GitHubSearchResponse struct {
		Items []*Repo
	}
	Repo struct {
		Fullname string    `json:"full_name"`
		Stars    int       `json:"stargazers_count"`
		PushedAt time.Time `json:"pushed_at"` // HLrecent
		Owner    struct {  // HLrecent
			Login       string  // HLrecent
			RecentRepos []*Repo // HLrecent
		} // HLrecent
	}
)

func main() {
	sr, err := searchPopularByLanguage("python", "go")
	checkErr(err)
	for _, r := range sr.Items {
		recent, err := recentUserRepos(r.Owner.Login) // HLrecent
		checkErr(err)
		r.Owner.RecentRepos = recent.Items // HLrecent
	}
	printResponse(os.Stdout, sr)
}

func searchPopularByLanguage(query, language string) (*GitHubSearchResponse, error) {
	queryString := `q=` + query + `+language:` + language + `&sort=stars&order=desc&per_page=3`
	return searchRepo(queryString)
}

func recentUserRepos(user string) (*GitHubSearchResponse, error) {
	queryString := `q=user:` + user + `&sort=updated&per_page=3`
	return searchRepo(queryString)
}

func searchRepo(queryString string) (*GitHubSearchResponse, error) {
	var sr *GitHubSearchResponse
	url := `https://api.github.com/search/repositories?` + queryString
	requestCounter++
	log.Printf("[%d] HTTP GET %s\n", requestCounter, url)
	resp, err := http.Get(url)
	log.Printf("[%d] HTTP RESPONSE\n", requestCounter)
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

func printResponse(w io.Writer, sr *GitHubSearchResponse) {
	now := time.Now()
	width := 30
	fmt.Fprintln(w, " #  Repo URL")
	fmt.Fprintln(w, strings.Repeat("-", width+35))
	for i, r := range sr.Items {
		fmt.Fprintf(w, "%02d. %-*s stars: %5d  pushed: %3.0fd ago\n",
			i+1, width, r.Fullname, r.Stars, daysSince(now, r.PushedAt))
		for _, rr := range r.Owner.RecentRepos { // HLrecent
			fmt.Fprintf(w, "    %-*s        %5d          %3.0fd ago\n", // HLrecent
				width, rr.Fullname, rr.Stars, daysSince(now, rr.PushedAt)) // HLrecent
		} // HLrecent
		fmt.Fprintf(w, "    %s\n", strings.Repeat("-", width+31)) // HLrecent
	}
}

func daysSince(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Hours() / 24
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
