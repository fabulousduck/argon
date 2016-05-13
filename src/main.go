package main

import (
	"fmt"
	"gopkg.in/readline.v1"
)

func main() {
	repl()
}

func repl() {
	p := NewPrsr()
	rl, err := readline.New("rocket> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		gline, err := rl.Readline()
		if err != nil {
			panic(err)
		}
		r1 := p.Run(gline)
		if err != nil {
			panic(err)
			continue
		}
		fmt.Println(r1, " memes")
	}
}
