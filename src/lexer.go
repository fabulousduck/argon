package main

import (
	"io/ioutil"
	"fmt"
	// "argon/src/scanner"
)



func main() {
	currentLine  			:= 0
	currentColl  			:= 0
	sourceIndex  			:= 0
	indexing     			:= false
	indexerStart 			:= 0
	tracker      			:= 0
	const_string_mode := false;
	int_count_mode    := false;
	symb_count_mode   := false;
	cookieJar         := []string{}
	file, _           := ioutil.ReadFile("../testfiles/main.ar")
	lexedToken        := &token{"0",0,0,0,"0"}

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

					cookieJar         = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
					tracker           = 0;
					indexerStart      = 0;
					indexing          = false;
					const_string_mode = false;

				}

	//if its tracking but not tracking a litteral string
		}else if(indexing && !const_string_mode){

			cookieJar    = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
			tracker      = 0;
			indexerStart = 0;
			indexing     = false;

			}else if(lexedToken.tokenType == "INTERGER"  && !const_string_mode){
				fmt.Println("found int")
				if(!int_count_mode){
					indexerStart = lexedToken.sourceIndex;
				}
				int_count_mode = true;
				tracker++;

			}else if(int_count_mode && lexedToken.tokenType != "INTERGER"){

				cookieJar      = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
				int_count_mode = false;
				indexerStart   = 0;
				tracker  		   = 0;

		 }else if(!const_string_mode && lexedToken.tokenType == "SYMBOL" ){
			//  fmt.Println("found symbol : ", lexedToken.cargo)
			 if(!symb_count_mode){
				 	indexerStart = lexedToken.sourceIndex;
			 }
			 symb_count_mode = true;
			 tracker++

		 }else if(!const_string_mode && symb_count_mode && lexedToken.tokenType != "SYMBOL"){

			 cookieJar      = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))
			 symb_count_mode = false;
			 indexerStart    = 0;
			 tracker  		   = 0;
		 }


		if char.cargo == "NEWLINE" {
			currentLine += 1
			currentColl = 0
		}
		sourceIndex += 1
		currentColl += 1

	}





//test. remove once we confirm it works

			fmt.Println(cookieJar)

}


func lex(char *char) *token {

	//if the char is a letter
  if(isIn(char.cargo, IDENTIFIER_STARTCHARS())){

		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"STRING"}

	}else if(isIn(char.cargo, NUMBER_CHARS())){

		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"INTERGER"}

	}else if(isIn(char.cargo, STRING_CHARACTERS())){

			return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"STRING_CONSTANT"}

	}else if(flavourType(char.cargo, ONE_CHARACTER_SYMBOLS())){
			fmt.Println("got : ", char.cargo);
			return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"SYMBOL"}

	}
		return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"ERROR"}


}

//function for checking if a char is a given type
func isIn(character string, section []string) bool{
  for i := 0; i < len(section); i++ {
    if section[i] == character{
    return true
		}
	}
	return false
}

func flavourType(character string, ttype []cookie) bool {

	for i := 0; i < len(ttype); i++{
		if ttype[i].cargo == character{
			return true
		}
	}
	return false
}

//function for fetching all the chars from a found string
func StackCookies(r []int) []string{
	cookieStack := []string{}

	for i := 0; i < r[1]; i++ {

		cookieStack = append(cookieStack, GET("../testfiles/main.ar", r[0]+i, r[3], r[0]+i).cargo);
	}
	return cookieStack;
}

// func flavourCookie(token *token) *cookie {
//
// }


//function to hang a type to concatenated or constant strings;
// func validate(toValidate string) {
//
// }

//this function is for the parser to request tokens from the lexer.
// func eat(tokenList []string, TkIndex int) *token{
//
// }



//function for making strings out of chars in the cookie stack
func concatCookie(r []string) string {
	f := "";

	for i := 0; i < len(r); i++ {

		f += r[i];

	}

	return f;
}
