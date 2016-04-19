package main

import(
  "fmt"
)

type ast_fn struct {
  node_name string
  node_args []string
  node_return string
  node_body string
}

type ast_if struct {
	node_statement []cookie{}
	node_body []string}{}
	node_end
	node_start
}


type initializer struct {
  ast_begin
  ast_beginln
  ast_end
  ast_endln
}


func accept(token cookie) bool {

}

func expect(token cookie) bool {

}

func throwError(errtype string, errToken string) {

}

