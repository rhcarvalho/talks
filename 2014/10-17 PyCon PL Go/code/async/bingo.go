package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

var (
	cardSize = flag.Int("card", 20, "how many numbers in a card")
	maxNum   = flag.Int("max", 100, "biggest number in a card + 1")
	winNum   = flag.Int("win", 10, "how many numbers a player needs to hear to win")
	pace     = flag.Duration("pace", 750*time.Millisecond, "duration before next number is drawn")
	speak    = flag.Bool("speak", true, "speak numbers as they're drawn")
)

const (
	// Reset
	ColorOff = "\033[0m" // Text Reset
	// Regular Colors
	Black  = "\033[0;30m" // Black
	Red    = "\033[0;31m" // Red
	Green  = "\033[0;32m" // Green
	Yellow = "\033[0;33m" // Yellow
	Blue   = "\033[0;34m" // Blue
	Purple = "\033[0;35m" // Purple
	Cyan   = "\033[0;36m" // Cyan
	White  = "\033[0;37m" // White
)

var colors = []string{Red, Green, Blue, Yellow, Purple, Cyan}

type Player struct {
	Name      string
	card      []int
	hits      int
	broadcast chan int
}

func NewPlayer(name string) *Player {
	p := &Player{
		Name:      name,
		broadcast: make(chan int)}
	p.card = rand.Perm(*maxNum)[:*cardSize]
	// for i := 0; i < *cardSize; i++ {
	// 	p.card = append(p.card, i)
	// }
	return p
}

func (p *Player) Play(bingo, chat chan<- string) {
	for n := range p.broadcast {
		hit := false
		for _, v := range p.card {
			if v == n {
				hit = true
				p.hits++
				break
			}
		}
		if p.hits >= *winNum {
			time.Sleep(time.Duration(rand.Float64() * float64(*pace+1*time.Millisecond)))
			bingo <- "bingo!!!"
		} else if hit {
			go func(hits int) {
				if hits == *winNum-1 {
					chat <- "almost there!"
				} else {
					chat <- "got it"
				}
			}(p.hits)
		}
	}
}

type Game struct {
	players []*Player
	numbers []int
	pos     int
	bingo   chan string
	chat    chan string
}

func NewGame(names ...string) *Game {
	g := &Game{
		numbers: rand.Perm(*maxNum),
		bingo:   make(chan string),
		chat:    make(chan string)}
	for _, name := range names {
		g.players = append(g.players, NewPlayer(name))
	}
	return g
}

func (g *Game) Play() {
	for i, p := range g.players {
		i, p := i, p
		b, c := make(chan string), make(chan string)
		go p.Play(b, c)
		go func() {
			for m := range b {
				g.bingo <- fmt.Sprintf("%s%s says: %s%s", colors[i%len(colors)], p.Name, m, ColorOff)
			}
		}()
		go func() {
			for m := range c {
				g.chat <- fmt.Sprintf("%s%s says: %s%s", colors[i%len(colors)], p.Name, m, ColorOff)
			}
		}()
	}
	numbers := make(chan int)
	go func() {
		for _, n := range g.numbers {
			numbers <- n
		}
		close(numbers)
	}()
	for {
		select {
		case b := <-g.bingo:
			done := make(chan struct{})
			go func() {
				exec.Command("say", "binnnnnng goooool!").Run()
				done <- struct{}{}
			}()
			log.Println(b)
			<-done
			return
		case c := <-g.chat:
			log.Println(c + "\007")
		case <-time.After(*pace):
			n, ok := <-numbers
			if !ok {
				log.Println("No winner :(")
				return
			}
			g.Broadcast(n)
		}
	}

}

func (g *Game) Broadcast(n int) {
	log.Printf("Number %d\n", n)
	if *speak {
		exec.Command("say", strconv.Itoa(n)).Run()
	}
	for _, p := range g.players {
		go func(p *Player) {
			p.broadcast <- n
		}(p)
	}
}

func (g *Game) next() int {
	n := g.numbers[g.pos]
	g.pos++
	return n
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	g := NewGame("Paweł", "Piotr", "Rodolfo", "Mariusz")

	fmt.Println("binGo v0.1")
	fmt.Println("Pace:", *pace)
	fmt.Println("Card size:", *cardSize)
	fmt.Println("Max number:", *maxNum)
	fmt.Println("Hits to win:", *winNum)
	fmt.Println("Tonight playing with us:")
	for i, p := range g.players {
		fmt.Printf("%d. %s%s%s %v\n", i+1, colors[i%len(colors)], p.Name, ColorOff, p.card)
	}
	fmt.Println()

	g.Play()
}
