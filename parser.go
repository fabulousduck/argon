package rocket

import(
	"fmt"
)

var precedanceTable = map[string]Operator{
	"STAR": {3, "LEFT"},
	"SLASH": {3, "LEFT"},
	"PLUS": {2, "LEFT"},
	"MIN": {2, "LEFT"},
}

func NewParser() *Parser {
	return &Parser{
		units: UnitTable{},
	}
}


func (p *Parser) Run(ln Program) int {
	units := lex(ln)
	p.units = units;
	p.parse()
	return 0
}


func (p *Parser) parse() int {
	outputQue := UnitTable{}
	operatorStack := UnitTable{}
	
	//evaluate to postfix notation
	for i := 0; i < len(p.units); i++ {
		currentUnit := p.units[i]

		switch currentUnit.tokenType {
			case "INTERGER":
				fmt.Println("interger case")
				outputQue = append(outputQue, currentUnit)
				break
			case "OPERATOR":
				fmt.Println("operator case")
				if(len(operatorStack) >= 1) {
					fmt.Println("there is an operator on the operator stack : ", operatorStack)
					for len(operatorStack) != 0 {
						if operatorStack.top().hasHigherPrecedance(currentUnit) == true {
							fmt.Println("operator ", operatorStack.top(), "has higher prec than", currentUnit)
							outputQue = append(outputQue, operatorStack.top())
							fmt.Println("appended to opq : ", outputQue)
							fmt.Println("popping from operator stack : ", operatorStack)
							operatorStack.pop()
						} else {
							break
						}
						
					}
				}
				fmt.Println("pushing ", currentUnit, "to the operator stack")
				operatorStack = append(operatorStack, currentUnit)
				break
		}
	}
	for len(operatorStack) != 0 {
		outputQue = append(outputQue, operatorStack[len(operatorStack)-1])
		operatorStack.pop()
	}

	fmt.Println("outstack ", outputQue)
	return 0
	//evaluate postix to outcome
}

// func exec(right cookie, op string, left cookie) cookie {
// 	rhs, _ := strconv.Atoi(right.cargo)
// 	lhs, _ := strconv.Atoi(left.cargo)
// 	switch op {
// 	case "+":
// 		return cookie{strconv.Itoa(lhs + rhs), "INTERGER", 0}
// 	case "-":
// 		return cookie{strconv.Itoa(lhs - rhs), "INTERGER", 0}
// 	case "*":
// 		return cookie{strconv.Itoa(lhs * rhs), "INTERGER", 0}
// 	case "/":
// 		return cookie{strconv.Itoa(lhs / rhs), "INTERGER", 0}
// 	}

// 	return unit{"0", "0", 0}
// }

func (stack UnitTable) top() Unit {
	if(len(stack) < 1) {
		panic("top of empty stack")
	}
	return stack[len(stack)-1]
}

func (s *UnitTable) pop() {
	*s = (*s)[:len(*s)-1]

}

func (unitA Unit) hasHigherPrecedance(unitB Unit) bool {
	operatorA := precedanceTable[unitA.notation]
	operatorB := precedanceTable[unitB.notation]

	return (operatorA.association == "LEFT" && operatorB.precedence <= operatorA.precedence)|| (operatorA.association == "RIGHT" && operatorB.precedence > operatorA.precedence)
}
