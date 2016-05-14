package rocket

/*
	List of keyword defenitions
*/
func keyWords() []defCookie {
	return []defCookie{
		{"if", "TOKEN_IF"},
		{"then", "TOKEN_THEN"},
		{"else", "TOKEN_ELSE"},
		{"while", "TOKEN_WHILE"},
		{"elif", "TOKEN_ELIF"},
		{"endif", "TOKEN_ENDIF"},
		{"while", "TOKEN_WHILE"},
		{"for", "TOKEN_FOR"},
		{"print", "IDENTIFIER_PRINT"},
		{"return", "TOKEN_RETURN"},
		{"exit", "TOKEN_EXIT"},
		{"define", "TOKEN_DEFINE"},
		{"fn", "TOKEN_FUNCTION"},
		{"memes", "TOKEN_MEMES"},
		{"int", "IDENTIFIER_INT"},
		{"uint", "IDENTFIER_INT_U"},
		{"string", "IDENTIFIER_STRING"},
		{"char", "IDENTIFIER_CHAR"},
		{"bool", "IDENTIFIER_BOOL"},
		{"double", "IDENTIFIER_DOUBLE"},
		{"float", "IDENTIFIER_FLOAT"},
		{"pingu", "TOKEN_PINGU"},
	}
}

/*
	List of one character long symbols
*/
func oneCharacterSymbols() []defCookie {
	return []defCookie{
		{"=", "SYMB_EQUALS"},
		{"(", "SYMB_LFT_PAR"},
		{")", "SYMB_RGT_PAR"},
		{"<", "SYMB_LFT"},
		{">", "SYMB_RGT"},
		{"/", "SYMB_SLASH"},
		{"*", "SYMB_STAR"},
		{"+", "SYMB_PLUS"},
		{"-", "SYMB_MIN"},
		{"!", "SYMB_BANG"},
		{"&", "SYMB_AND"},
		{".", "SYMB_DOT"},
		{",", "SYMB_COMMA"},
		{";", "SYMB_END_OF_STATEMENT"},
		{"{", "SYMB_LFT_CB"},
		{"}", "SYMB_RGT_CB"},
	}
}

/*
	List of two character long symbols
*/
func twoCharacterSymbols() []defCookie {
	return []defCookie{
		{"==", "SYMB_IS_EQ"},
		{"<=", "SYMB_LFT_EQ"},
		{">=", "SYMB_RGT_EQ"},
		{"!=", "SYMB_ISNOT_EQ"},
		{"++", "SYMB_PLUS_PLUS"},
		{"**", "SYMB_STAR_STAR"},
		{"--", "SYMB_MIN_MIN"},
		{"+=", "SYMB_PLUS_EQ"},
		{"-=", "SYMB_MIN_EQ"},
		{"||", "SYMB_OR"},
	}
}

func letters() []string {
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

func digits() []string {
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

func KEYWORDS() []defCookie {
	return keyWords()
}

func IDENTIFIER_STARTCHARS() []string {
	return letters()
}

func IDENTIFIER_CHARS() []string {
	return letters()
}

func NUMBER_STARTCHARS() []string {
	return digits()
}

func NUMBER_CHARS() []string {
	return digits()
}

func STRING_CHARACTERS() []string {
	return []string{"''", "\""}
}

func WHITESPACE_CHARS() []string {
	return []string{"\\n", "\\t"}
}

func ONE_CHARACTER_SYMBOLS() []defCookie {
	return oneCharacterSymbols()
}

func TWO_CHARACTER_SYMBOLS() []defCookie {
	return twoCharacterSymbols()
}
