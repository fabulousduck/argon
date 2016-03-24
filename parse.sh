#! /bin/sh
declare -a files
files[0]="scanner.go"
files[1]="lexer.go"
files[2]="token.go"
files[3]="symbols.go"
$preDir="~/Desktop/Golang/src/argon/src"
gp $predir/files[0] $predir/files[1] $predir/files[2] $predir/files[3]

