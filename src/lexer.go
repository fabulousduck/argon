package main

import (
	"fmt"
	"io/ioutil"
	// "argon/src/scanner"
)



func main() {
	currentLine  := 0
	currentColl  := 0
	sourceIndex  := 0
	indexing     := false
	indexerStart := 0
	tracker      := 0
	cookieJar    := []string{}
	file, err    := ioutil.ReadFile("../testfiles/main.ar")
	lexedToken   := &token{"0",0,0,0,"0"}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("char sourceIndex lineIndex collumnIndex")
	for i := 0; i < len(string(file)); i++ {
		char := GET("../testFiles/main.ar", currentColl, currentLine, sourceIndex);



    /*
      do fany lexer stuff here
    */
    lexedToken = lex(char);

		if lexedToken.tokenType == "STRING" {
			if(!indexing){
				indexerStart = lexedToken.sourceIndex;
			}

			indexing = true;
			tracker++;
			fmt.Println("oh ! a cookie at : " , lexedToken.sourceIndex);
			fmt.Println(lexedToken.cargo);

		}else if(indexing && lexedToken.tokenType != "STRING"){

			fmt.Println("oh. no more cookies. cookie trail lasted  : ", indexerStart , "/" , indexerStart + tracker);
			cookieJar = append(cookieJar, concatCookie(StackCookies([]int{indexerStart,tracker,indexerStart+tracker-1,lexedToken.lineIndex})))

			tracker = 0;
			indexerStart = 0;
			indexing = false;
			fmt.Println(cookieJar);
		}

		if char.cargo == "NEWLINE" {
			currentLine += 1
			currentColl = 0
		}
		sourceIndex += 1
		currentColl += 1

		}
    /*
    end fancy lexer stuff here
    */

	}

//this function is for the parser to request tokens from the lexer.
// func eat(tokenList []string, TkIndex int) *token{
//
// }

func lex(char *char) *token {

	//if the char is a letter
  if(isIn(char.cargo, IDENTIFIER_STARTCHARS())){

    return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex,"STRING"}
  }else{
		// fmt.Println(buffer);
	}
	// else if(isIn(char.cargo, )){
	//
	// }
	return &token{char.cargo,char.sourceIndex,char.lineIndex,char.colIndex, "ERROR"}
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
	}
	return cookieStack;
}

func concatCookie(r []string) string {
	f := "";
	for i := 0; i < len(r); i++ {
		f += r[i];
	}
	return f;
}
