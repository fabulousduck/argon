package main

import (
	"fmt"
	"github/fabulousduck/rocket"
	"gopkg.in/readline.v1"
)

func main() {
	repl()
}

func repl() {
	p := rocket.NewPrsr()
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
