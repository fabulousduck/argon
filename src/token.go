package main

type token struct {
  cargo string
  sourceIndex, lineIndex, colIndex int
  tokenType string
}

type cookie struct {
  cargo string
  t_sort string
  line_no int
}

type defCookie struct {
  cargo string
  t_sort string
}
