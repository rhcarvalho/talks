package main

import (
	"fmt"
	"time"
)

type Pythonista struct {
	Name, City string
	Since      time.Time
	Happy      bool
}

type MyInt int

type ReverseString string

// -types OMIT

func (p *Pythonista) String() string { // HL
	var happy string
	if p.Happy {
		happy = "Happy"
	} else {
		happy = "Mostly happy"
	}
	return fmt.Sprintf("%v, from %v.\n%v Pythonista for %v.",
		p.Name, p.City, happy, time.Since(p.Since))
}

func (i MyInt) Add(other MyInt) MyInt { // HL
	return i + other
}

func (rs ReverseString) String() string { // HL
	runes := []rune(string(rs))
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// -methods OMIT

func main() {
	since := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(&Pythonista{"Rodolfo", "Rio de Janeiro", since, true})

	var x, y MyInt = 6, 5
	fmt.Println(x.Add(y).Add(8))

	s := "Python 编程语言"
	fmt.Println(s)
	fmt.Println(ReverseString(s))
	fmt.Println(ReverseString(s) + "another string")
}
