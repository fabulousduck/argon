package main

import(
  "fmt"
  "io/ioutil"
)

func scanner_driver(){
  currentLine := 0
  currentColl := 0
  sourceIndex := 0
  file, _ := ioutil.ReadFile("../testFiles/main.ar")
  fmt.Println("char sourceIndex lineIndex collumnIndex")
  for i := 0; i < len(string(file)); i++ {
    char := get("../testFiles/main.ar", currentColl, currentLine, sourceIndex)
    fmt.Println(char.cargo, char.sourceIndex, char.lineIndex, char.colIndex)
    if char.cargo == "NEWLINE" {
      currentLine += 1
      currentColl = 0
    }
    sourceIndex += 1
    currentColl += 1

  }

}
