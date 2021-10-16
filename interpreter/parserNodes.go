package interpreter

type IntNode struct {
    Type string
    Value int
}

type VarAcessNode struct {
    Type string 
    Name string
}

type StringNode struct {
    Type string 
    Value string
}

type EmptyNode struct {
    Type string
}
