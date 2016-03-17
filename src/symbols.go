package main

/*
	List of keyword defenitions
*/
func keyWords() []string{
	return []string{
		"if",
		"then",
		"else",
		"while",
		"elif",
		"endif",
		"while",
		"loop",
		"endloop",
		"print",
		"return",
		"exit",
		"define",
	}
}
	/*
		List of one character long symbols
	*/
func oneCharacterSymbols() []string{
	return []string{
		"=",
		"(",
		")",
		"<",
		">",
		"/",
		"*",
		"+",
		"-",
		"!",
		"&",
		".",
		",",
		";",
	}
}
	/*
		List of two character long symbols
	*/
func twoCharacterSymbols() []string{
	return []string{
		"==",
		"<=",
		">=",
		"<>",
		"!=",
		"++",
		"**",
		"--",
		"+=",
		"-=",
		"||",
	}
}

func letters() []string{
	return []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
}

func digits() []string{
	return []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
}



func IDENTIFIER_STARTCHARS() []string{
	return letters()
}

func IDENTIFIER_CHARS() []string{
	return letters()
}

func NUMBER_STARTCHARS() []string{
	return digits();
}

func NUMBER_CHARS() []string{
	return digits()
}

func STRING_CHARACTERS() []string{
	return []string{"''","\""}
}

func WHITESPACE_CHARS() []string{
	return []string{"\\n","\\t"}
}
