package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "github.com/leoDesilva/lisp-interpreter/lexer"
)

func ReadFile(filename string) string{
    filePointer, _ := os.Open(filename) 
    fileBytes, _ := ioutil.ReadAll(filePointer)
    return string(fileBytes)
}

func main(){
    filename := os.Args[1]
    file := ReadFile(filename)
    fmt.Println(file)

    lexer := Lexer{file,string(file[0]),0}
    fmt.Println(lexer)

    fmt.Println(IsLexer)
}
