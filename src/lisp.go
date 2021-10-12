package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "github.com/leoDesilva/lisp-interpreter/interpreter"
    "strings"
)

func ReadFile(filename string) string{
    filePointer, _ := os.Open(filename) 
    fileBytes, _ := ioutil.ReadAll(filePointer)
    return string(fileBytes)
}

func main(){
    type Lex = interpreter.Lexer

    filename := os.Args[1]
    file := ReadFile(filename)
    formattedFile := strings.Replace(file, `\n`, ``, -1)

    lexer := Lex{
        File: formattedFile,
        Char: string(formattedFile[0]),
        Index: 0,
    }

    fmt.Println(lexer.File)

    tokens := interpreter.Lex(&lexer)
    fmt.Println(tokens)
}
