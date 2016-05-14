package main

import (
	"fmt"
	"gopkg.in/readline.v1"
	"github.com/fabulousduck/rocket"
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
	r1 := 0
	for {
		gline, err := rl.Readline()
		if err != nil {
			panic(err)
		}
		r1 = p.Run(gline)
		fmt.Println(r1)
	}

}
