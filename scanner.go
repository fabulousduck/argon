package rocket

type char struct {
	cargo       string
	sourceIndex int
	lineIndex   int
	colIndex    int
}

func GetChar(file string, coll int, line int, sourceIndex int) *char {
	wantedChar := string(file[sourceIndex])
	if wantedChar == " " {
		wantedChar = "SPACE"
	}
	if wantedChar == "\n" {
		wantedChar = ""
	}

	return &char{wantedChar, sourceIndex, line, coll}

}
