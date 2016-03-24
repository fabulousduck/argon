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
		{"fn", "TOKEN_FUNCTION"},
		{"memes","TOKEN_MEMES"},
		{"int","TOKEN_INT"},
		{"int32", "TOKEN_INT_32"},
		{"int64", "TOKEN_INT_64"},
		{"uint", "TOKEN_INT_U"},
		{"uint32", "TOKEN_INT_U_32"},
		{"uint64", "TOKEN_INT_U_64"},
		{"string", "TOKEN_STRING"},
		{"char", "TOKEN_CHAR"},
		{"bool", "TOKEN_BOOL"},
		{"double","TOKEN_DOUBLE"},
		{"float","TOKEN_FLOAT"},

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
		{";", "TOKEN_END_OF_STATEMENT"},
		{"{", "TOKEN_LFT_CB"},
		{"}", "TOKEN_RGT_CB"},
	}
}
	/*
		List of two character long symbols
	*/
func twoCharacterSymbols() []cookie{
	return []cookie{
		{"==", "TOKEN_IS_EQ"},
		{"<=", "TOKEN_LFT_EQ"},
		{">=", "TOKEN_RGT_EQ"},
		{"!=", "TOKEN_ISNOT_EQ"},
		{"++", "TOKEN_PLUS_PLUS"},
		{"**", "TOKEN_STAR_STAR"},
		{"--", "TOKEN_MIN_MIN"},
		{"+=", "TOKEN_PLUS_EQ"},
		{"-=", "TOKEN_MIN_EQ"},
		{"||", "TOKEN_OR"},
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

func ONE_CHARACTER_SYMBOLS() []cookie{
	return oneCharacterSymbols()
}

func TWO_CHARACTER_SYMBOLS() []cookie{
	return twoCharacterSymbols()
}
