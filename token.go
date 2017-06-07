package rocket

type association int

type token struct {
	cargo                            string
	sourceIndex, lineIndex, colIndex int
	tokenType                        string
}

type cookie struct {
	cargo   string
	t_sort  string
	line_no int
}

type unit struct {
	cargo  string
	notation string
	tokenType string
}

type Parser struct {
	tokens []cookie
	pos    int
	tok    cookie
}

type ops struct {
	prec  int
	assoc string
}
