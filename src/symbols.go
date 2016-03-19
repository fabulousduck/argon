package main

/*
	List of keyword defenitions
*/
func keyWords() []cookie{
	return []cookie{
		{"if",   "TOKEN_IF"},
		{"then", "TOKEN_THEN"},
		{"else", "TOKEN_ELSE"},
		{"while","TOKEN_WHILE"},
		{"elif", "TOKEN_ELIF"},
		{"endif","TOKEN_ENDIF"},
		{"while","TOKEN_WHILE"},
		{"for",  "TOKEN_FOR"},
		{"print","TOKEN_PRINT"},
		{"return","TOKEN_RETURN"},
		{"exit",  "TOKEN_EXIT"},
		{"define","TOKEN_DEFINE"},
	}
}
	/*
		List of one character long symbols
	*/
func oneCharacterSymbols() []cookie{
	return []cookie{
		{"=", "TOKEN_EQUALS"},
		{"(", "TOKEN_LFT_PAR"},
		{")", "TOKEN_RGT_PAR"},
		{"<", "TOKEN_LFT"},
		{">", "TOKEN_RGT"},
		{"/", "TOKEN_SLASH"},
		{"*", "TOKEN_STAR"},
		{"+", "TOKEN_PLUS"},
		{"-", "TOKEN_MIN"},
		{"!", "TOKEN_BANG"},
		{"&", "TOKEN_AND"},
		{".", "TOKEN_DOT"},
		{",", "TOKEN_COMMA"},
		{";", "TOKEN_ENDOFSTATEMENT"},
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

func ONE_CHARACTER_SYMBOLS() []string{
	return oneCharacterSymbols()
}

func TWO_CHARACTER_SYMBOLS() []string{
	return twoCharacterSymbols()
}
