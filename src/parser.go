package main

import (
	"fmt"
)

func main() {
	//returns a tokenSet
	//the tokens are stored in structs called cookies
	tokenSet := lexicallyAnalize("../testfiles/main.ar")

	fmt.Println(tokenSet)
}

// func checkStatment(token cookie){
//   switch token.t_sort :
//     case "TOKEN_IF":
//       parse_if()
// }

// func parse_if(){}
//
// func parse_fn(){
//
// }
//
// func parse_expr(){
//
// }
