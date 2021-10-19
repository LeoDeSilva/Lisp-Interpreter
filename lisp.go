package main

import (
    //"fmt"
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
    type Parse = interpreter.Parser
    type Interpret = interpreter.Interpreter

    filename := os.Args[1]
    file := ReadFile(filename)
    formattedFile := strings.Replace(file, `\n`, ``, -1)

    lexer := Lex{
        File: formattedFile,
        Char: string(formattedFile[0]),
        Index: 0,
    }

    tokens := interpreter.Lex(&lexer)

    parser := Parse{
        Tokens: tokens,
        Token: tokens[0],
        Index: 0,
    }

    ast, err := interpreter.Parse(&parser)
    if err {return}

    evaluator := Interpret{
        AST: ast,
        Node: ast[0],
        Index: 0,
        Variables: make(map[string]interface{}),
        Functions: make(map[string]interface{}),
    }

    err = interpreter.Interpret(&evaluator)
    if err {return}
}
