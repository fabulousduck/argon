package rocket

import (
	//"fmt"
	"strconv"
	"strings"
)

type Stack []cookie

var opers = map[string]ops{
	"STAR":  {6, "LEFT"},
	"PLUS":  {5, "LEFT"},
	"MIN":   {5, "LEFT"},
	"SLASH": {6, "LEFT"},
}

func NewParser() *Parser {
	return &Parser{
		pos: 0,
	}
}

func (p *Parser) parse(rl string) int {

	//fmt.Println("p.tokens : ", p.tokens)
	//shunting yard algorrithm
	var (
		stack  Stack
		output Stack
	) //numbers
	count := 0

	for i := 0; i < len(p.tokens); i++ {

		switch {

		case p.tokens[count].isEmpty():

			break

		case p.tokens[count].isNumber():

			output = append(output, p.tokens[count])
			count++
		case p.tokens[count].isOperator():

			if len(stack) == 0 {
				stack = append(stack, p.tokens[count])
				count++
				break
			}
			if len(stack) >= 1 {
				o1 := p.tokens[count]
				o2 := stack.Top()

				//move the top of the stack to the ouput
				if o1.isHigherPrec(o2) {
					t := output.Top()
					output.Pop()
					e := exec(t, stack.Top().cargo, output.Top())
					stack.Pop()
					output.Pop()
					output = append(output, e)
				}

			}
			stack = append(stack, p.tokens[count])

			count++
		case p.tokens[count].isComma():
			//loop this untill the top of the stack is a comma
			for {
				output = append(output, stack[len(stack)-1])
				stack = append(stack[:len(stack)-1], stack[len(stack)+1:]...)
			}
			count++

		case p.tokens[count].isLPAREN():
			//fmt.Println("found left paren")
			stack = append(stack, p.tokens[count])
			count++

		case p.tokens[count].isRPAREN():

			for stack.Top().cargo != "(" {
				fg := output.Top()
				output.Pop()
				f := exec(fg, stack.Top().cargo, output.Top())
				stack.Pop()
				output.Pop()
				output = append(output, f)
			}
			if stack.Top().cargo == "(" {

			}
			count++
		}

	}

	if len(stack) != 0 {
		//fmt.Println("stack not empty, found  : ", stack)
		ls := 0
		for len(stack) != 0 {
			t := output[len(output)-1]
			output.Pop()
			ls = len(stack) - 1
			e := exec(t, stack[ls].cargo, output[len(output)-1])
			stack.Pop()
			output.Pop()
			output = append(output, e)

			//also exec here
		}
	}
	r, _ := strconv.Atoi(output[0].cargo)
	return r
}

func exec(right cookie, op string, left cookie) cookie {
	rhs, _ := strconv.Atoi(right.cargo)
	lhs, _ := strconv.Atoi(left.cargo)
	switch op {
	case "+":
		return cookie{strconv.Itoa(lhs + rhs), "INTERGER", 0}
	case "-":
		return cookie{strconv.Itoa(lhs - rhs), "INTERGER", 0}
	case "*":
		return cookie{strconv.Itoa(lhs * rhs), "INTERGER", 0}
	case "/":
		return cookie{strconv.Itoa(lhs / rhs), "INTERGER", 0}

	}

	return cookie{"0", "0", 0}
}

func (s *Stack) PopTo(dest *Stack) {
	hold := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	*dest = append(*dest, hold)
}

func (s *Stack) Top() cookie {

	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (p *Parser) Run(ln string) int {
	tokens := lexicallyAnalize(ln)

	for i := 0; i < len(tokens); i++ {
		if tokens[i].t_sort == "empty" {
			tokens = append(tokens[:i], tokens[i+1:]...)
		}
	}

	p.flush()
	p.tokens = tokens

	return p.parse(ln)
}

func (c *cookie) isRPAREN() bool {
	return c.cargo == ")"
}

func (c *cookie) isLPAREN() bool {
	return c.cargo == "("
}

func (c *cookie) isNumber() bool {
	return c.t_sort == "INTERGER"
}

func (c *cookie) isOperator() bool {
	if c.t_sort == "empty" {
		return false
	}
	tok := strings.Split(c.t_sort, "_")[1]
	var ok bool
	if _, ok := opers[tok]; ok {

		return ok
	}
	return ok
}

func (c *cookie) isComma() bool {
	return strings.Split(c.t_sort, "_")[1] == "COMMA"
}

func (c *cookie) isEmpty() bool {
	return c.t_sort == "empty"
}

func (p *Parser) flush() {
	p.pos = 0
	p.tokens = nil
}

func (o1 cookie) isHigherPrec(o2 cookie) bool {
	o1o := strings.Split(o1.t_sort, "_")[1]
	o2o := strings.Split(o2.t_sort, "_")[1]

	if opers[o1o].assoc == "LEFT" && opers[o2o].prec >= opers[o1o].prec ||
		opers[o1o].assoc == "RIGHT" && opers[o2o].prec > opers[o1o].prec {
		return true
	}
	return false
}
