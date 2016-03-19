package main

import (
	"io/ioutil"
	"fmt"
	// "argon/src/scanner"
)



func main() {
	currentLine  := 0
	currentColl  := 0
	sourceIndex  := 0
	indexing     := false
	indexerStart := 0
	tracker      := 0
	const_string_mode := false;
	cookieJar    := []string{}
	file, err    := ioutil.ReadFile("../testfiles/main.ar")
	lexedToken   := &token{"0",0,0,0,"0"}
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(string(file)); i++ {

		char := GET("../testFiles/main.ar", currentColl, currentLine, sourceIndex);

    lexedToken = lex(char);


		//this handles reading out non-string constants

		// if its a string or litteral string token and its not tracking a litteral string
		if lexedToken.tokenType == "STRING" || lexedToken.tokenType == "STRING_CONSTANT" && !const_string_mode{
			//if its not tracking
			if(!indexing){
				indexerStart = lexedToken.sourceIndex;
			}
			// if its a litteral string token
			if(lexedToken.tokenType == "STRING_CONSTANT"){
				const_string_mode = true;
				fmt.Println("found a \" ")
			}

			indexing = true;
			tracker++;
			//if tracking a litteral string
		}else if(indexing && const_string_mode){

			tracker++
				//if it finds a litteral string token while already lexing a litteral string
				if(lexedToken.tokenType == "STRING_CONSTANT"){

					cookieJar = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
					fmt.Println("indexing and const_string_mode and tokenType == STRING_CONSTANT")
					tracker = 0;
					indexerStart = 0;
					indexing = false;
					const_string_mode = false;

				}
	//if its tracking but not tracking a litteral string
	}else if(indexing && !const_string_mode){

		cookieJar = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
		fmt.Println("tracker for append = ", tracker);
		tracker = 0;
		indexerStart = 0;
		indexing = false;

	}








		if char.cargo == "NEWLINE" {
			currentLine += 1
			currentColl = 0
		}

		sourceIndex += 1
		currentColl += 1

		}
		//test. remove once we confirm it works
		for i := 0; i < len(cookieJar); i++ {
			fmt.Println(cookieJar[i])
		}
	}

//this function is for the parser to request tokens from the lexer.
// func eat(tokenList []string, TkIndex int) *token{
//
// }

func lex(char *char) *token {

	//if the char is a letter
  if(isIn(char.cargo, IDENTIFIER_STARTCHARS())){

		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"STRING"}

	}else if(isIn(char.cargo, NUMBER_CHARS())){

		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"INTEGER"}

	}else if(isIn(char.cargo, STRING_CHARACTERS())){

			return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"STRING_CONSTANT"}

	}else if(isIn(char.cargo, ONE_CHARACTER_SYMBOLS())){

			return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"SYMBOL"}

	}
		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"ERROR"}


}

func isIn(character string, section []string) bool{
  for i := 0; i < len(section); i++ {
    if section[i] == character{
    return true
		}
	}
	return false
}


func StackCookies(r []int) []string{
	cookieStack := []string{}

	for i := 0; i < r[1]; i++ {

		//cookies in the cookieJar are listed as follows : indexerStart /
		//																								 tracker			/
		//																					indexerStart + tracker -1 /
		//																								line 					/
		cookieStack = append(cookieStack, GET("../testfiles/main.ar", r[0]+i, r[3], r[0]+i).cargo);
				fmt.Println("adding : ", GET("../testfiles/main.ar", r[0]+i, r[3], r[0]+i).cargo, "to stack. stack is now" , cookieStack);
	}
	fmt.Println("SC output : " , cookieStack);
	return cookieStack;
}

//function to hang a type to concatenated or constant strings;
// func validate(toValidate string) {
//
// }

//function for concatinating cookies
func concatCookie(r []string) string {
	f := "";
	for i := 0; i < len(r); i++ {
		f += r[i];
	}
	// fmt.Println("outputted : ", f );
	return f;
}
