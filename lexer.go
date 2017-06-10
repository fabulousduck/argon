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
		fmt.Println(unit, unitType, currentIndex)
		switch unitType {
			case "INTERGER":
				nextType := determineType(program.peek(currentIndex+1))
				if nextType == "INTERGER" {
					completeUnit := program.peekLong(currentIndex, "INTERGER")
					unitList = append(unitList, completeUnit)
					currentIndex += len(completeUnit.cargo)			
				} else {
					unitList = append(unitList, Digits.getBaseUnit(unit))
					currentIndex++
				}
			case "LETTER": //not needed for now
			case "OPERATOR":
				nextUnit := program.peek(currentIndex+1)
				nextType := determineType(program.peek(currentIndex+1))
				if nextType == "OPERATOR" {
					newOperator := unit+nextUnit
					definition := TwoCharacterSymbols.getBaseUnit(newOperator)
					if  definition.cargo != "BASE_LOOKUP_ERR" {
						unitList = append(unitList,definition)
						currentIndex += 2
					} else {
						unitList = append(unitList,OneCharacterSymbols.getBaseUnit(unit))
						currentIndex ++
					}
				} else {
					unitList = append(unitList,OneCharacterSymbols.getBaseUnit(unit))
					currentIndex++
				}
			case "SYMBOL":
				unitList = append(unitList, OneCharacterSymbols.getBaseUnit(unit))
				currentIndex++
		}
	}
	fmt.Println(unitList)
	return unitList
}

func (program Program) peek (index int) string {
	return string(program[index])
}

func (program Program) peekLong (startIndex int, unitType string) Unit {
	currentIndex := startIndex
	accum := ""
	for ;; {
		if(currentIndex >= len(program)) {
			return Unit{cargo: accum, notation: unitType, unitType: unitType}
			break
		}
		currentUnit := string(program[currentIndex])
		if determineType(currentUnit) != unitType {
			return Unit{cargo: accum, notation: unitType, unitType: unitType}
			break
		} else {
			accum += currentUnit
			currentIndex++
		}
	}

	return Unit{cargo: "PEEKLONG_ERR", notation: "PEEKLONG_ERR", unitType: "INTERGER"}
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
	return Unit{cargo: "BASE_LOOKUP_ERR", notation: "BASE_LOOKUP_ERR", unitType: "BASE_LOOKUP_ERR"}
}

func (unitTable UnitTable) getType (currentUnit string) string {
	for _, tableUnit := range unitTable {
		if tableUnit.cargo == currentUnit {
            return tableUnit.unitType
        }
	}
	return "UNKNOWN_TYPE"
}