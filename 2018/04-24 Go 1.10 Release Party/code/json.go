//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Rename struct {
	Old string
	New string
}

func main() {
	dec := json.NewDecoder(strings.NewReader(`
		{"old": "big", "new": "large"}
		{"old": "yesterday", "new": "tomorrow"}
		{"old": "sun", "new": "rain", "comment": "REVIEW"}
	`))
	dec.DisallowUnknownFields()
	for dec.More() {
		var r Rename
		err := dec.Decode(&r)
		fmt.Println(r, err)
	}
}
