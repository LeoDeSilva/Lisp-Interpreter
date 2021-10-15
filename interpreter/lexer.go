package interpreter 

import (
    "fmt"
    "strings"
)

var LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
var DIGITS = "0123456789"

var TT_LPAREN = "TT_LPAREN"
var TT_RPAREN = "TT_RPAREN"
var TT_ADD = "TT_ADD"
var TT_SUB = "TT_SUB"
var TT_MUL = "TT_MUL" 
var TT_DIV = "TT_DIV"

var TT_EQ = "TT_EQ"
var TT_EE = "TT_EE"
var TT_NE = "TT_NE"
var TT_GT = "TT_GT"
var TT_GTE = "TT_GTE"
var TT_LT = "TT_LT"
var TT_LTE = "TT_LTE"
var TT_NOT = "TT_NOT"

var TT_KEYWORD = "TT_KEYWORD"
var TT_IDENTIFIER = "TT_IDENTIFIER"
var TT_INT = "TT_INT"
var TT_STRING = "TT_STRING"

type Lexer struct {
    File string 
    Char string 
    Index int
}

type Node struct {
    Type string
    Value string
}


func printNode(n Node){
    fmt.Println(n.Type, n.Value)
}


func advance(l *Lexer){
    l.Index++
    if l.Index < len(l.File) - 1 {
        l.Char = string(l.File[l.Index]) 
    }
}


func retreat(l *Lexer){
    l.Index--
    l.Char = string(l.File[l.Index]) 
}


func lexString(l *Lexer) (string, bool) {
    var stringValue string

    for l.Index < len(l.File) - 1 {
        if l.Char == `"` {
            return stringValue, false
        }

        stringValue = stringValue + l.Char
        advance(l)
    }

    return "", true
}


func lexIdentifier(l *Lexer) (string, bool){
    var identifier string 

    for l.Index < len(l.File) - 1 {
        if strings.Contains(LETTERS, l.Char) {
            identifier = identifier + l.Char
        } else{ 
            retreat(l)
            return identifier, false
        }
        advance(l)
    }

    return "", true
}


func lexNumber(l *Lexer) (string, bool){
    var number string 

    for l.Index < len(l.File) - 1 {
        if strings.Contains(DIGITS, l.Char) {
            number = number + l.Char
        } else{
            retreat(l)
            return number, false
        }
        advance(l)
    }

    return "", true
}


func lexDouble(l *Lexer) (Node, bool){
    if string(l.File[l.Index + 1]) == "=" {
        char := l.Char
        advance(l)
        if char == "=" { 
            return Node{TT_EE, "=="}, false 
        } else if char == ">" { 
            return Node{TT_GTE, ">="}, false 
        } else if char == "<" { 
            return Node{TT_LTE, ">="}, false 
        } else if char == "!" {
            return Node{TT_NE, "!="}, false
        }
    } else{
        if l.Char == "=" { 
            return Node{TT_EQ, "="}, false
        } else if l.Char == ">" { 
            return Node{TT_GT, ">"}, false
        } else if l.Char == "<" { 
            return Node{TT_LT, ">"}, false
        } else if l.Char == "!" {
            return Node{TT_NOT, "!"}, false
        }
    }

    return Node{"",""}, true
}


func lexToken(l *Lexer) (Node, bool) {
    if l.Char == "(" {
        return Node{TT_LPAREN,"("}, false 
    } else if l.Char == ")" {
        return Node{TT_RPAREN,")"}, false 
    } else if l.Char == "+" {
        return Node{TT_ADD, "+"}, false
    } else if l.Char == "-" {
        return Node{TT_SUB, "-"}, false
    } else if l.Char == "*" {
        return Node{TT_MUL, "*"}, false
    } else if l.Char == "/" {
        return Node{TT_DIV, "/"}, false
    } else if l.Char == "="{
        node, err := lexDouble(l)
        if (err) { return Node{"",""}, true }
        return node, false
    } else if l.Char == ">" {
        node, err := lexDouble(l)
        if (err) { return Node{"",""}, true }
        return node, false
    } else if l.Char == "<" {
        node, err := lexDouble(l)
        if (err) { return Node{"",""}, true }
        return node, false
    } else if l.Char == "!"{
        node, err := lexDouble(l)
        if (err) { return Node{"",""}, true }
        return node, false
    } else if l.Char == `"` {
        advance(l)
        stringValue, err := lexString(l)
        if (err) { return Node{"",""}, true }
        return Node{TT_STRING, stringValue}, false
    } else if strings.Contains(LETTERS, l.Char) {
        identifier, err := lexIdentifier(l)
        if(err) { return Node{"",""}, true }
        return Node{TT_IDENTIFIER, identifier}, false
    } else if strings.Contains(DIGITS, l.Char) {
        number, err := lexNumber(l)
        if(err) { return Node{"",""}, true } 
        return Node{TT_INT,number}, false
    }

    return Node{"",""}, true
}


func Lex(l *Lexer) []Node{
    var nodes []Node
    for l.Index < len(l.File) - 1{
        node,err := lexToken(l) 
        if (!err){
            //printNode(node)
            nodes = append(nodes, node)
        }
        advance(l) 
    }
    return nodes
}
