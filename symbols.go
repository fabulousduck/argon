package rocket

/*
	List of keyword defenitions
*/
var Keywords = []unit{
	{"int", "INT", "IDENTIFIER"},
	{"uint", "UINT", "IDENTFIER"},
	{"string", "STRING", "IDENTIFIER"},
	{"char", "CHAR", "IDENTIFIER"},
	{"bool", "BOOLEAN", "IDENTIFIER"},
	{"double", "DOUBLE", "IDENTIFIER"},
	{"float", "FLOAT", "IDENTIFIER"},
}

/*
	List of one character long symbols
*/
var OneCharacterSymbols = []unit{
	{"=", "EQUALS", "OPERATOR"},
	{"(", "LEFT_PARENTHESES", "SYMBOL"},
	{")", "RIGHT_PARENTHESES", "SYMBOL"},
	{"<", "LEFT", "OPERATOR"},
	{">", "RIGHT", "OPERATOR"},
	{"/", "SLASH", "OPERATOR"},
	{"*", "STAR", "OPERATOR"},
	{"+", "PLUS", "OPERATOR"},
	{"-", "MIN", "OPERATOR"},
	{"!", "BANG", "OPERATOR"},
	{"&", "AND", "OPERATOR"},
	{".", "DOT", "SYMBOL"},
	{",", "COMMA", "SYMBOL"},
	{";", "END_OF_STATEMENT", "SYMBOL"},
	{"{", "LEFT_CURLY_BRACKET", "SYMBOL"},
	{"}", "RIGHT_CURLY_BRACKET", "SYMBOL"},
	{"'", "SINGLE_QUOTE", "SYMBOL"},
	{"\"", "DOUBLE_QUOTE", "SYMBOL"},

}

/*
	List of two character long symbols
*/
var TwoCharacterSymbols = []unit{
	{"==", "IS_EQUAL", "OPERATOR"},
	{"<=", "LEFT_EQUAL", "OPERATOR"},
	{">=", "RIGHT_EQUAL", "OPERATOR"},
	{"!=", "ISNOT_EQUAL", "OPERATOR"},
	{"++", "PLUS_PLUS", "OPERATOR"},
	{"**", "STAR_STAR", "OPERATOR"},
	{"--", "MIN_MIN", "OPERATOR"},
	{"+=", "PLUS_EQUALS", "OPERATOR"},
	{"-=", "MIN_EQUALS", "OPERATOR"},
	{"||", "OR", "OPERATOR"},
	{"\\t", "TAB", "OPERATOR"},
	{"\\n", "NEWLINE", "OPERATOR"},
}


//made these units for the sake of consistency
var Letters = []unit{
	{"a", "A", "LETTER"}
	{"b", "B", "LETTER"}
	{"c", "C", "LETTER"}
	{"d", "D", "LETTER"}
	{"e", "E", "LETTER"}
	{"f", "F", "LETTER"}
	{"g", "G", "LETTER"}
	{"h", "H", "LETTER"}
	{"i", "I", "LETTER"}
	{"j", "J", "LETTER"}
	{"k", "K", "LETTER"}
	{"l", "L", "LETTER"}
	{"m", "M", "LETTER"}
	{"n", "N", "LETTER"}
	{"o", "O", "LETTER"}
	{"p", "P", "LETTER"}
	{"q", "Q", "LETTER"}
	{"r", "R", "LETTER"}
	{"s", "S", "LETTER"}
	{"t", "T", "LETTER"}
	{"u", "U", "LETTER"}
	{"v", "V", "LETTER"}
	{"w", "W", "LETTER"}
	{"x", "X", "LETTER"}
	{"y", "Y", "LETTER"}
	{"z", "Z", "LETTER"}
}

var Digits = []unit{
	{"0" "ZERO", "INTEGER"},
	{"1" "ONE", "INTEGER"},
	{"2" "TWO", "INTEGER"},
	{"3" "THREE", "INTEGER"},
	{"4" "FOUR", "INTEGER"},
	{"5" "FIVE", "INTEGER"},
	{"6" "SIX", "INTEGER"},
	{"7" "SEVEN", "INTEGER"},
	{"8" "EIGHT", "INTEGER"},
	{"9" "NINE", "INTEGER"},
}