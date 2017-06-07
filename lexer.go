package rocket

import(
	"unicode"
)

func lex(program string) []token {
	for i := 0; i < len(program); i++ {

	}
}

func determineType(item string) string {
	if(unicode.IsDigit(item)) {
		return "integer"
	}
	if(contains)
}

func (unitTable UnitTable) contains(unitTable []unit, currentUnit string) bool {
    for _, tableUnit := range unitTable {
        if tableUnit == currentUnit {
            return true
        }
    }
    return false
}