package main

import (
	"fmt"
	"github.com/fabulousduck/rocket"
	// "gopkg.in/readline.v1"
)

func main() {
	var test rocket.Program = "1*277/(3+4-45)"
	p := rocket.NewParser()
	result := p.Run(test)
	fmt.Println(result)

}
