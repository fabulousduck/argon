package rocket

import (
	"fmt"
)

type Program string

func lex (program Program) []Unit {

	currentIndex := 0;
	unitList := []Unit{}

	for currentIndex < len(program) {
		unit := string(program[currentIndex])
		unitType := determineType(unit)
		
		switch unitType {
			case "INTERGER":
				nextType := program.peek(currentIndex+1)
				if nextType == "INTERGER" {
					completeUnit := program.peekLong(currentIndex, "INTERGER")
					unitList = append(unitList, completeUnit)
					currentIndex += len(completeUnit.cargo)				
				} else {
					unitList = append(unitList, Digits.getBaseUnit(unit))
					currentIndex++
				}
			case "LETTER":
			case "OPERATOR":
			case "SYMBOL":
		}
		currentIndex++
	}
	fmt.Println(unitList)
	return unitList
}

func (program Program) peek (index int) string {
	unit := string(program[index])
	return determineType(unit)
}

func (program Program) peekLong (startIndex int, unitType string) Unit {
	currentIndex := startIndex
	accum := ""
	for ;; {
		if(currentIndex >= len(program)) {
			return Unit{cargo: accum, notation: unitType, tokenType: unitType}
			break
		}
		currentUnit := string(program[currentIndex])
		if determineType(currentUnit) != unitType {
			return Unit{cargo: accum, notation: unitType, tokenType: unitType}
			break
		} else {
			accum += currentUnit
			currentIndex++
		}
	}

	return Unit{cargo: "PEEKLONG_ERR", notation: "PEEKLONG_ERR", tokenType: "INTERGER"}
}

func determineType (unit string) string {
	if Digits.contains(unit) {
		return "INTERGER"
	}
	if Letters.contains(unit) {
		return "LETTER"
	}
	if OneCharacterSymbols.contains(unit) {
		return OneCharacterSymbols.getType(unit)
	}
	return "UNDEFINED_TYPE"
}

func (unitTable UnitTable) contains (currentUnit string) bool {
    for _, tableUnit := range unitTable {
        if tableUnit.cargo == currentUnit {
            return true
        }
    }
    return false
}

func (unitTable UnitTable) getBaseUnit (currentUnit string) Unit {
	for _, tableUnit := range unitTable {
		if tableUnit.cargo == currentUnit {
            return tableUnit
        }
	}
	return Unit{cargo: "PEEKLONG_ERR", notation: "PEEKLONG_ERR", tokenType: "INTERGER"}
}

func (unitTable UnitTable) getType (currentUnit string) string {
	for _, tableUnit := range unitTable {
		if tableUnit.cargo == currentUnit {
            return tableUnit.tokenType
        }
	}
	return "UNKNOWN_TYPE"
}