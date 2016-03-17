package main

import (
	"fmt"
	"io/ioutil"
)

type char struct {
	cargo       string
	sourceIndex int
	lineIndex   int
	colIndex    int
}


func GET(file string, coll int, line int, sourceIndex int) *char {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("could not find file")
	}

	stringFile := string(f)
	wantedChar := string(stringFile[sourceIndex])
	if wantedChar == " " {
		wantedChar = "SPACE"
	}
	if wantedChar == "\n" {
		wantedChar = "NEWLINE"
	}

	return &char{wantedChar, sourceIndex, line, coll}

}
