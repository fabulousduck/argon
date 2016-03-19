package main

type token struct {
  cargo string
  sourceIndex int
  lineIndex int
  colIndex int
  tokenType string
}

type cookie struct {
  cargo string
  t_type string
}
