# Rocket,  An expression parser written in go
[![forthebadge](http://forthebadge.com/images/badges/designed-in-ms-paint.svg)](http://forthebadge.com)
[![forthebadge](http://forthebadge.com/images/badges/powered-by-jeffs-keyboard.svg)](http://forthebadge.com)

## usage

Rocket is still very much being built and is not ready for proper use yet.
The current version should only be used for testing perposes.


## dependencies

- Go
- version 1.6

## near future todo's

1. Lexer
  * Make function for combining char and int for types like int32/64 and uint32/64
2. Parser
  * implement shunting-yard algorithm

## How to install.
 **Windows.**
    Currently there is no command/.bat for compiling on windows.
    This is high on our priority list tho.
    For now you will have to compile allthe files manually buy going into the src folder and executing the following:
    ```
      go run lexer.go scanner.go symbols.go token.go
    ```
  
 **Linux/OSX**
    For mac and linux users. we have the build.sh file.
    Since this is also still quite buggy. it is reccomended that you use the same way windows user currently compile it
