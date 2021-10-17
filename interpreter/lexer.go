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

var TT_EOF = "TT_EOF"

var TT_VAR_ACCESS = "TT_VAR_ACCESS"
var TT_BIN_OP = "TT_BIN_OP"
var TT_VAR_ASSIGN = "TT_VAR_ASSIGN"
var TT_FUNCTION_CALL = "TT_FUNCTION_CALL"
var TT_BLOCK = "TT_BLOCK"
var TT_FUNCTION_DEFENITION = "TT_FUNCTION_DEFENITION"
var TT_IF = "TT_IF"

var KEYWORDS = map[string]bool{
    "setf":true,
    "block":true,
}

type Lexer struct {
    File string 
    Char string 
    Index int
}

type Node struct {
    Type string
    Value string
    Class string
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

func matches(t Node, value string) bool {
    return t.Value == value
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
            return Node{TT_EE, "==", "BIN_OP"}, false 
        } else if char == ">" { 
            return Node{TT_GTE, ">=", "BIN_OP"}, false 
        } else if char == "<" { 
            return Node{TT_LTE, ">=", "BIN_OP"}, false 
        } else if char == "!" {
            return Node{TT_NE, "!=", "BIN_OP"}, false
        }
    } else{
        if l.Char == "=" { 
            return Node{TT_EQ, "=", "EQ"}, false
        } else if l.Char == ">" { 
            return Node{TT_GT, ">", "BIN_OP"}, false
        } else if l.Char == "<" { 
            return Node{TT_LT, ">", "BIN_OP"}, false
        } else if l.Char == "!" {
            return Node{TT_NOT, "!", "UNARY_OP"}, false
        }
    }

    return Node{"","","ERROR"}, true
}


func lexToken(l *Lexer) (Node, bool) {
    if l.Char == "(" {
        return Node{TT_LPAREN,"(","BRACKET"}, false 
    } else if l.Char == ")" {
        return Node{TT_RPAREN,")","BRACKET"}, false 
    } else if l.Char == "+" {
        return Node{TT_ADD, "+", "BIN_OP"}, false
    } else if l.Char == "-" {
        return Node{TT_SUB, "-", "BIN_OP"}, false
    } else if l.Char == "*" {
        return Node{TT_MUL, "*", "BIN_OP"}, false
    } else if l.Char == "/" {
        return Node{TT_DIV, "/", "BIN_OP"}, false
    } else if l.Char == "="{
        node, err := lexDouble(l)
        if (err) { return Node{"","","ERROR"}, true }
        return node, false
    } else if l.Char == ">" {
        node, err := lexDouble(l)
        if (err) { return Node{"","","ERROR"}, true }
        return node, false
    } else if l.Char == "<" {
        node, err := lexDouble(l)
        if (err) { return Node{"","","ERROR"}, true }
        return node, false
    } else if l.Char == "!"{
        node, err := lexDouble(l)
        if (err) { return Node{"","","ERROR"}, true }
        return node, false
    } else if l.Char == `"` {
        advance(l)
        stringValue, err := lexString(l)
        if (err) { return Node{"","","ERROR"}, true }
        return Node{TT_STRING, stringValue, "ATOM"}, false

    } else if strings.Contains(LETTERS, l.Char) {
        identifier, err := lexIdentifier(l)
        if(err) { return Node{"","","ERROR"}, true }
        if KEYWORDS[identifier] {
            return Node{TT_KEYWORD, identifier, "KEYWORD"}, false
        }
        return Node{TT_IDENTIFIER, identifier, "ATOM"}, false

    } else if strings.Contains(DIGITS, l.Char) {
        number, err := lexNumber(l)
        if(err) { return Node{"","","ERROR"}, true } 
        return Node{TT_INT,number, "ATOM"}, false
    }

    return Node{"","","ERROR"}, true
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
    nodes = append(nodes, Node{TT_EOF, "","END"})
    return nodes
}
