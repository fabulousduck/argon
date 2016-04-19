#! /bin/sh
preDir="src"
scanDir=${preDir}/scanner.go
lexrDir=${preDir}/lexer.go
toknDir=${preDir}/token.go
symbDir=${preDir}/symbols.go
parsDir=${preDir}/parser.go

go run $scanDir $lexrDir $toknDir $symbDir $parsDir
