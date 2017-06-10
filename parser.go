package rocket

import(
	"fmt"
	"strconv"
)

var precedanceTable = map[string]Operator{
	"STAR": {3, "LEFT"},
	"SLASH": {3, "LEFT"},
	"PLUS": {2, "LEFT"},
	"MIN": {2, "LEFT"},
	"RIGHT_PARENTHESES": {110, "LEFT"},
	"LEFT_PARENTHESES": {110, "LEFT"},
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


func (p *Parser) toPostFix() UnitTable {
	outputQue := UnitTable{}
	operatorStack := UnitTable{}
	
	//evaluate to postfix notation
	for i := 0; i < len(p.units); i++ {
		currentUnit := p.units[i]
		fmt.Println("operatorStack : ", operatorStack)
		fmt.Println("currently P   : ", currentUnit.cargo)
		switch currentUnit.unitType {
			case "INTERGER":
				outputQue = append(outputQue, currentUnit)
				break
			case "OPERATOR":
				if len(operatorStack) >= 1 && operatorStack.parenIsNotTop() {
					for len(operatorStack) != 0 {
						if operatorStack.top().hasHigherPrecedance(currentUnit) &&  operatorStack.top().unitType != "SYMBOL"{
							outputQue = append(outputQue, operatorStack.top())
							operatorStack.pop()
						} else {
							break
						}
						
					}
				}
				operatorStack = append(operatorStack, currentUnit)
				break
			case "SYMBOL": //find a way to refacor this
				switch currentUnit.symbolType() {
					case "LEFT_PARENTHESES":
						operatorStack = append(operatorStack, currentUnit)
						fmt.Println("LFT_PARENT OPSTACK : ", operatorStack)
						break
					case "RIGHT_PARENTHESES":
						if(operatorStack.top().symbolType() != "LEFT_PARENTHESES") {
							for operatorStack.top().symbolType() != "LEFT_PARENTHESES" {
								outputQue = append(outputQue, operatorStack.top())
								operatorStack.pop()
							}
						} else {
							operatorStack.pop()
							operatorStack.pop()
							break
						}
				}
				break
		}
	}
	for len(operatorStack) != 0 {
		outputQue = append(outputQue, operatorStack[len(operatorStack)-1])
		operatorStack.pop()
	}
	fmt.Println("opq : ", outputQue)
	return outputQue
}

func (stack UnitTable) parenIsNotTop() bool {
	return stack[len(stack)-1].unitType != "SYMBOL"
}

func (unit Unit) symbolType() string {
	v := OneCharacterSymbols.getBaseUnit(unit.cargo).notation
	return v
}

func postFixToOutcome (postFix UnitTable) int {
	stack := UnitTable{}

	for i := 0; i < len(postFix); i++ {
		currentUnit := postFix[i]

		switch currentUnit.unitType {
			case "INTERGER":
				stack = append(stack, currentUnit)
				break
			case "OPERATOR":
				temp := stack.top()
				stack.pop()
				res := exec(temp, currentUnit, stack.top())
				stack.pop()
				stack = append(stack, res)
				break
		}
	}
	fmt.Println(stack)
	result, _ :=strconv.Atoi(stack.top().cargo)
	return result
} 

func (p *Parser) parse() int {
	postFix := p.toPostFix()
	fmt.Println(postFix)
	return postFixToOutcome(postFix)
	//evaluate postix to outcome
}

func exec(right Unit, operator Unit, left Unit) Unit {
	rhs, _ := strconv.Atoi(right.cargo)
	lhs, _ := strconv.Atoi(left.cargo)
	switch operator.cargo {
	case "+":
		return Unit{strconv.Itoa(lhs + rhs), "INTERGER", "INTERGER"}
	case "-":
		return Unit{strconv.Itoa(lhs - rhs), "INTERGER", "INTERGER"}
	case "*":
		return Unit{strconv.Itoa(lhs * rhs), "INTERGER", "INTERGER"}
	case "/":
		return Unit{strconv.Itoa(lhs / rhs), "INTERGER", "INTERGER"}
	}

	return Unit{"0", "0", "0"}
}

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
