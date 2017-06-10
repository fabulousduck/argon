package rocket

/*
	List of keyword defenitions
*/

var Keywords = UnitTable{
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
var OneCharacterSymbols = UnitTable{
	{"=", "EQUALS", "OPERATOR"},
	{"(", "LEFT_PARENTHESES", "SYMBOL"},
	{")", "RIGHT_PARENTHESES", "SYMBOL"},
	{"<", "LEFT", "OPERATOR"},
	{">", "RIGHT", "OPERATOR"},
	{"^", "POWER", "OPERATOR"},	
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
var TwoCharacterSymbols = UnitTable{
	{"==", "IS_EQUAL", "OPERATOR"},
	{"<=", "LEFT_EQUAL", "OPERATOR"},
	{">=", "RIGHT_EQUAL", "OPERATOR"},
	{"!=", "ISNOT_EQUAL", "OPERATOR"},
	{"++", "PLUS_PLUS", "OPERATOR"},
	{"**", "STAR_STAR", "OPERATOR"},
	{"--", "MIN_MIN", "OPERATOR"},
	{"+=", "PLUS_EQUALS", "OPERATOR"},
	{"*=", "STAR_EQUALS", "OPERATOR"},
	{"/=", "DEVIDE_EQUALS", "OPERATOR"},
	{"-=", "MIN_EQUALS", "OPERATOR"},
	{"||", "OR", "OPERATOR"},
	{"\\t", "TAB", "OPERATOR"},
	{"\\n", "NEWLINE", "OPERATOR"},
}


//made these units for the sake of consistency
var Letters = UnitTable{
	{"a", "A", "LETTER"},
	{"b", "B", "LETTER"},
	{"c", "C", "LETTER"},
	{"d", "D", "LETTER"},
	{"e", "E", "LETTER"},
	{"f", "F", "LETTER"},
	{"g", "G", "LETTER"},
	{"h", "H", "LETTER"},
	{"i", "I", "LETTER"},
	{"j", "J", "LETTER"},
	{"k", "K", "LETTER"},
	{"l", "L", "LETTER"},
	{"m", "M", "LETTER"},
	{"n", "N", "LETTER"},
	{"o", "O", "LETTER"},
	{"p", "P", "LETTER"},
	{"q", "Q", "LETTER"},
	{"r", "R", "LETTER"},
	{"s", "S", "LETTER"},
	{"t", "T", "LETTER"},
	{"u", "U", "LETTER"},
	{"v", "V", "LETTER"},
	{"w", "W", "LETTER"},
	{"x", "X", "LETTER"},
	{"y", "Y", "LETTER"},
	{"z", "Z", "LETTER"},
}

var Digits = UnitTable{
	{"0", "ZERO", "INTERGER"},
	{"1", "ONE", "INTERGER"},
	{"2", "TWO", "INTERGER"},
	{"3", "THREE", "INTERGER"},
	{"4", "FOUR", "INTERGER"},
	{"5", "FIVE", "INTERGER"},
	{"6", "SIX", "INTERGER"},
	{"7", "SEVEN", "INTERGER"},
	{"8", "EIGHT", "INTERGER"},
	{"9", "NINE", "INTERGER"},
}