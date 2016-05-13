package main

import (
	"fmt"
)

func lexicallyAnalize(f string) []cookie {
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

	cookieJar := []cookie{}
	lexedToken := &token{"0", 0, 0, 0, "0"}
	eof := false
	fmt.Println(string(f))
	for i := 0; i < len(f); i++ {
		char := GET(string(f), currentColl, currentLine, sourceIndex)
		lexedToken = lex(char)
		if i == len(string(f))-1 {
			eof = true
		}

		//paste old code back here if needid//

		//paste old code back here if needid//

		switch lexedToken.tokenType {
		case "STRING":

			if !string_indexing && !const_indexing {
				string_indexing = true
				string_index_start = lexedToken.sourceIndex
			}
			if string_indexing && !const_indexing {
				string_tracker++
			}

			if const_indexing {
				const_tracker++
			}
			if eof {
				string_tracker++

				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})), "STRING", lexedToken.lineIndex})
				string_indexing = false
				string_tracker = 0

			}

			if int_indexing && !eof && !const_indexing {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})), "INTERGER", lexedToken.lineIndex})
				int_indexing = false
				int_tracker = 0

			}

			if symb_indexing && !eof && !const_indexing {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})), "SYMBOL", lexedToken.lineIndex})
				symb_indexing = false
				symb_tracker = 0

			}

		case "STRING_CONSTANT":

			if int_indexing && !eof && !const_indexing {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})), "INTERGER", lexedToken.lineIndex})
				int_indexing = false
				int_tracker = 0

			}

			if symb_indexing && !eof && !const_indexing {

				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})), "SYMBOL", lexedToken.lineIndex})
				symb_indexing = false
				symb_tracker = 0

			}

			if const_indexing && !eof {
				const_tracker++
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{const_index_start, const_tracker, const_index_start + const_tracker - 1, lexedToken.lineIndex})), "STRING_CONSTANT", lexedToken.lineIndex})
				const_indexing = false
				const_tracker = 0

			}
			if !const_indexing {
				const_indexing = true
				const_index_start = lexedToken.sourceIndex
				const_tracker++
			}

			if eof {
				const_tracker++
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{const_index_start, const_tracker, const_index_start + const_tracker - 1, lexedToken.lineIndex})), "STRING_CONSTANT", lexedToken.lineIndex})
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
			}

			if int_indexing && !const_indexing {
				int_tracker++
			}

			if string_indexing && !const_indexing && !eof {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})), "STRING", lexedToken.lineIndex})
				string_indexing = false
				string_tracker = 0
			}
			if symb_indexing && !const_indexing {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{symb_index_start, symb_tracker, symb_index_start + symb_tracker + 1, lexedToken.lineIndex})), "SYMBOL", lexedToken.lineIndex})
				symb_indexing = false
				symb_tracker = 0
			}

			if eof {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})), "INTERGER", lexedToken.lineIndex})
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
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})), "INTERGER", lexedToken.lineIndex})
				int_indexing = false
				int_tracker = 0
			}
			if string_indexing && !eof {
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{string_index_start, string_tracker, string_index_start + string_tracker + 1, lexedToken.lineIndex})), "STRING", lexedToken.lineIndex})
				string_indexing = false
				string_tracker = 0
			}

			if eof {
				symb_tracker++
				cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{symb_index_start, symb_tracker, symb_index_start + (symb_tracker + 2), lexedToken.lineIndex})), "SYMBOL", lexedToken.lineIndex})
				symb_indexing = false
				symb_tracker = 0

			}

		}

		if char.cargo == "" {
			currentLine += 1
			currentColl = 0
			// if its a newline it doesnt do shit
			if eof {
				if symb_indexing {
					symb_tracker++
					cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{symb_index_start, symb_tracker, symb_index_start + (symb_tracker + 2), lexedToken.lineIndex})), "SYMBOL", lexedToken.lineIndex})
					symb_indexing = false
					symb_tracker = 0

				}
				if int_indexing {
					cookieJar = append(cookieJar, cookie{concatCookie(StackCookies(f, []int{int_index_start, int_tracker, int_index_start + int_tracker + 1, lexedToken.lineIndex})), "INTERGER", lexedToken.lineIndex})
					int_indexing = false
					int_tracker = 0
				}
			}
		}
		sourceIndex += 1
		currentColl += 1

	}

	return cookieBox(cookieJar)

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

func seek(needle string, haystack []defCookie) string {
	for i := 0; i < len(haystack); i++ {
		if haystack[i].cargo == needle {
			return haystack[i].t_sort
		}
	}
	return "NOT_FOUND"
}

func flavourType(character string, ttype []defCookie) bool {

	for i := 0; i < len(ttype); i++ {
		if ttype[i].cargo == character {
			return true
		}
	}
	return false
}

//function for fetching all the chars from a found string
func StackCookies(f string, r []int) []string {
	cookieStack := []string{}

	for i := 0; i < r[1]; i++ {

		cookieStack = append(cookieStack, GET(f, r[0]+i, r[3], r[0]+i).cargo)
	}
	return cookieStack
}

//function to hang a type to concatenated or constant strings;
func validate(toValidate cookie) (cookie, []cookie) {

	switch toValidate.t_sort {
	case "STRING":
		if seek(toValidate.cargo, KEYWORDS()) != "NOT_FOUND" {
			// fmt.Println(toValidate.cargo, KEYWORDS(), " MEMEMEMEEMMEMEMEEs")
			return cookie{toValidate.cargo, seek(toValidate.cargo, KEYWORDS()), toValidate.line_no}, []cookie{}
		} else {
			return cookie{toValidate.cargo, "VARIABLE", toValidate.line_no}, []cookie{}
		}
	case "STRING_CONSTANT":
		return cookie{toValidate.cargo, "STRING_CONSTANT", toValidate.line_no}, []cookie{}
	case "INTERGER":
		return cookie{toValidate.cargo, "INTERGER", toValidate.line_no}, []cookie{}
	case "SYMBOL":
		if seek(toValidate.cargo, ONE_CHARACTER_SYMBOLS()) != "NOT_FOUND" {
			return cookie{toValidate.cargo, seek(toValidate.cargo, ONE_CHARACTER_SYMBOLS()), toValidate.line_no}, []cookie{}
		} else if seek(toValidate.cargo, TWO_CHARACTER_SYMBOLS()) != "NOT_FOUND" {
			return cookie{toValidate.cargo, seek(toValidate.cargo, TWO_CHARACTER_SYMBOLS()), toValidate.line_no}, []cookie{}
		} else {
			symbol_bucket := []cookie{}
			for i := 0; i < len(toValidate.cargo); i++ {
				symbol_bucket = append(symbol_bucket, cookie{string(toValidate.cargo[i]), "SYMBOL", toValidate.line_no})
			}
			return cookie{"empty", "empty", toValidate.line_no}, symbol_bucket
		}
	}

	return cookie{"ERR", "ERR", toValidate.line_no}, []cookie{}
}

//this function is for the parser to request tokens from the lexer.
func cookieBox(tokenList []cookie) []cookie {
	cookiePackage := []cookie{}
	for i := 0; i < len(tokenList); i++ {
		currentToken, symbol_bucket := validate(tokenList[i])
		// fmt.Println(tokenList)
		if len(symbol_bucket) > 1 {
			//	fmt.Println(symbol_bucket)
			for j := 0; j < len(symbol_bucket); j++ {
				ct, _ := validate(cookie{symbol_bucket[j].cargo, "SYMBOL", symbol_bucket[j].line_no})
				cookiePackage = append(cookiePackage, ct)
			}
		}
		cookiePackage = append(cookiePackage, currentToken)
	}
	return cookiePackage
}

//function for making strings out of chars in the cookie stack
func concatCookie(r []string) string {
	f := ""

	for i := 0; i < len(r); i++ {

		f += r[i]

	}

	return f
}
