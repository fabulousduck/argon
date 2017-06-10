package rocket

/*
	TYPES
		UnitTable
			A collection of units used for lookup
			In a UnitTable usable units for the parser are defined
		location
			Block of data the unit contains about the
			physical location of the unit in the prgram
		
*/


type UnitTable []Unit


type Unit struct {
	cargo  string
	notation string
	tokenType string
}

type Parser struct {
	units UnitTable
}

type Operator struct {
	precedence  int
	association string
}
