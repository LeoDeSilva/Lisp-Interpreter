package interpreter 

import (
    "fmt"
    "strconv"
)

type Parser struct {
    Tokens []Node
    Token Node 
    Index int
}

func next(p *Parser){
    p.Index++
    if (p.Index >= len(p.Tokens)){
        p.Index = -1
        return
    }
    p.Token = p.Tokens[p.Index]
}

func Parse(p *Parser) []interface{}{
    var ast []interface{}
    fmt.Println(p.Tokens)

    for p.Index != -1 && p.Index < len(p.Tokens) {
        node, _ := parseExpr(p)
        fmt.Println(node)
        next(p)
    }

    return ast
}

func parseExpr(p *Parser) (interface{},bool) {
    var node interface{}

    if p.Token.Type == TT_LPAREN{

    } else{
        node,_ = parseAtom(p)
        return node, false
    }

    return node, true
}

func parseList(p *Parser){

}

func parseAtom(p *Parser)(interface{}, bool){
    var node interface{}

    if p.Token.Type == TT_EOF{
        return EmptyNode{TT_EOF}, false

    } else if p.Token.Type == TT_INT{
        value, _ := strconv.Atoi(p.Token.Value)
        return IntNode{
            Type: TT_INT,
            Value: value,
        }, false

    } else if p.Token.Type == TT_STRING{
        return StringNode{
            Type: TT_STRING,
            Value: p.Token.Value,
        }, false

    } else if p.Token.Type == TT_IDENTIFIER{
        return VarAcessNode {
            Type: TT_VAR_ACCESS,
            Name: p.Token.Value,
        }, false
    }

    return node, true
}

