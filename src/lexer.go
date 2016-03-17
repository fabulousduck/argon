package main

import (
	"fmt"
	"io/ioutil"
	// "argon/src/scanner"
)

// fmt.Println("//");
// fmt.Println(indexing);
// fmt.Println(lexedToken.tokenType);
// fmt.Println(lex(GET("../testFiles/main.ar", currentColl+1, currentLine, sourceIndex+1)).tokenType);
// fmt.Println("//");

func main() {
	currentLine  := 0
	currentColl  := 0
	sourceIndex  := 0
	indexing     := false
	indexerStart := 0
	tracker      := 0
	positions    := [][]int{}
	file, err    := ioutil.ReadFile("../testfiles/main.ar")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("char sourceIndex lineIndex collumnIndex")
	for i := 0; i < len(string(file)); i++ {
		char := GET("../testFiles/main.ar", currentColl, currentLine, sourceIndex);



    /*
      do fany lexer stuff here
    */
    lexedToken := lex(char);
		if(!indexing && lexedToken.tokenType == "STRING"){

			indexerStart = lexedToken.sourceIndex;
			indexing = true;
			fmt.Println("started indexing at : " , indexerStart , " origin : " , sourceIndex);

			if(lex(GET("../testFiles/main.ar",currentColl+1, currentLine, sourceIndex+1)).tokenType != "STRING"){ // <--- it wont enter this state. fix this.
				fmt.Println("memes");
				indexing = false;
				fmt.Println("tracker ended at : " , tracker, " end origin : " , sourceIndex);
				positions = append(positions, []int{indexerStart,tracker,currentLine, sourceIndex,currentColl});
				fmt.Println(concat(putStack(positions[len(positions)-1])));

				};
		}else if(indexing && lexedToken.tokenType == "STRING"){

			tracker++;
			fmt.Println("membes2")

			if(lex(GET("../testFiles/main.ar",currentColl+1, currentLine, sourceIndex+1)).tokenType != "STRING"){ // <--- it wont enter this state. fix this.
				fmt.Println("memes2");
				indexing = false;
				fmt.Println("tracker 2 ended at : " , tracker);
				positions = append(positions, []int{indexerStart,tracker,currentLine, sourceIndex,currentColl});
				fmt.Println(concat(putStack(positions[len(positions)-1])));

				};
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


func lex(char *char) *token {

	//if the char is a letter
  if(isIn(char.cargo, IDENTIFIER_STARTCHARS())){
		// fmt.Println("found : ", char);
    // buffer = append(buffer, char.cargo);
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


func putStack(r []int) []string{
	stack := []string{}
	j := 0;

	for i :=r[0]; i < r[0]+r[1]; i++ {
		ch := GET("../testFiles/main.ar",r[0+j], r[2], r[3]);
		stack = append(stack, ch.cargo);
		j++
	}
	return stack;
}

func concat(r []string) string {
	f := "";
	for i := 0; i < len(r); i++ {
		f += r[i];
	}
	return f;
}
