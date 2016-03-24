package main

import (
	"fmt"
	"io/ioutil"
	// "argon/src/scanner"
)

func main() {
	currentLine := 0
	currentColl := 0
	sourceIndex := 0

	string_indexing := false
	string_tracker := 0
	string_index_start := 0

	const_indexing := false
	const_tracker := 0
	const_index_start := 0

	int_indexing := false
	int_tracker := 0
	int_index_start := 0

	symb_indexing := false
	symb_tracker := 0
	symb_index_start := 0

	cookieJar := []string{}
	file, _ := ioutil.ReadFile("../testfiles/main.ar")
	lexedToken := &token{"0", 0, 0, 0, "0"}
	eof := false

	for i := 0; i < len(string(file)); i++ {

		char := GET("../testFiles/main.ar", currentColl, currentLine, sourceIndex)

		lexedToken = lex(char)

		if i == len(string(file))-2 {
			eof = true
		}

		//paste old code back here if needid//

		//paste old code back here if needid//

		switch lexedToken.tokenType {
		case "STRING":

			if !string_indexing && !const_indexing {
				string_indexing = true
				string_index_start = lexedToken.sourceIndex
				// string_tracker++
			}
			if string_indexing {
				string_tracker++
			}

			if eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})))
				string_indexing = false
				string_tracker = 0

			}

			if int_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{int_index_start, int_tracker, int_index_start + int_tracker - 1, lexedToken.lineIndex})))
				int_indexing = false
				int_tracker = 0

			}

			if symb_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})))
				symb_indexing = false
				symb_tracker = 0

			}

		case "STRING_CONSTANT":

			if const_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{const_index_start, const_tracker, const_index_start + const_tracker - 1, lexedToken.lineIndex})))
				const_indexing = false
				const_tracker = 0
			}
			if !const_indexing {
				const_indexing = true
				const_index_start = lexedToken.sourceIndex
			}

			if eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{const_index_start, const_tracker, const_index_start + const_tracker - 1, lexedToken.lineIndex})))
				const_indexing = false
				const_tracker = 0
			}

		case "INTERGER":

			if const_indexing {
				const_tracker++
			}

			if !int_indexing && !const_indexing {
				int_indexing = true
				int_index_start = lexedToken.sourceIndex
				// int_tracker++
			}

			if int_indexing && !const_indexing {
				int_tracker++
			}

			if string_indexing && !const_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})))
				string_indexing = false
				string_tracker = 0
			}
			if symb_indexing && !const_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})))
				symb_indexing = false
				symb_tracker = 0
			}

			if eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})))
				int_indexing = false
				int_tracker = 0
			}

		case "SYMBOL":

			if const_indexing {
				const_tracker++
			}

			if !symb_indexing {
				symb_indexing = true
				symb_index_start = lexedToken.sourceIndex
				// symb_tracker++
			}

			if symb_indexing {
				symb_tracker++
			}

			if int_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})))
				int_indexing = false
				int_tracker = 0
			}
			if string_indexing && !eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})))
				string_indexing = false
				string_tracker = 0
			}

			if eof {
				cookieJar = append(cookieJar, concatCookie(StackCookies([]int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})))
				symb_indexing = false
				symb_tracker = 0
			}

		}

		if char.cargo == "NEWLINE" {
			currentLine += 1
			currentColl = 0
		}
		sourceIndex += 1
		currentColl += 1

	}

	//test. remove once we confirm it works

	fmt.Println(cookieJar)

}

func lex(char *char) *token {

	//if the char is a letter
	if isIn(char.cargo, IDENTIFIER_STARTCHARS()) {

		return &token{char.cargo, char.sourceIndex, char.lineIndex, char.colIndex, "STRING"}

	} else if isIn(char.cargo, NUMBER_CHARS()) {

		return &token{char.cargo, char.sourceIndex, char.lineIndex, char.colIndex, "INTERGER"}

	} else if isIn(char.cargo, STRING_CHARACTERS()) {

		return &token{char.cargo, char.sourceIndex, char.lineIndex, char.colIndex, "STRING_CONSTANT"}

	} else if flavourType(char.cargo, ONE_CHARACTER_SYMBOLS()) {

		return &token{char.cargo, char.sourceIndex, char.lineIndex, char.colIndex, "SYMBOL"}

	}
	return &token{char.cargo, char.sourceIndex, char.lineIndex, char.colIndex, "ERROR"}

}

//function for checking if a char is a given type
func isIn(character string, section []string) bool {
	for i := 0; i < len(section); i++ {
		if section[i] == character {
			return true
		}
	}
	return false
}

func flavourType(character string, ttype []cookie) bool {

	for i := 0; i < len(ttype); i++ {
		if ttype[i].cargo == character {
			return true
		}
	}
	return false
}

//function for fetching all the chars from a found string
func StackCookies(r []int) []string {
	cookieStack := []string{}

	for i := 0; i < r[1]; i++ {

		cookieStack = append(cookieStack, GET("../testfiles/main.ar", r[0]+i, r[3], r[0]+i).cargo)
	}
	return cookieStack
}

// func flavourCookie(token *token) *cookie {
//
// }

//function to hang a type to concatenated or constant strings;
// func validate(toValidate string) {
//
// }

//this function is for the parser to request tokens from the lexer.
// func eat(tokenList []string, TkIndex int) *token{
//
// }

//function for making strings out of chars in the cookie stack
func concatCookie(r []string) string {
	f := ""

	for i := 0; i < len(r); i++ {

		f += r[i]

	}

	return f
}
