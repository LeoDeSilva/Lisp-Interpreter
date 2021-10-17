package interpreter

type WhileNode struct {
    Type string 
    Condition interface{}
    Consequence interface{}
}


type IfNode struct {
    Type string 
    Condition interface{}
    Consequence interface{}
    Alternative interface{}
}

type FunctionDefenitionNode struct {
    Type string 
    Identifier string 
    Parameters []interface{} 
    Block interface{}
}


type BlockNode struct {
    Type string 
    Block []interface{}
}

type FunctionCallNode struct {
    Type string
    Identifier string
    Parameters []interface{}
}

type VarAssignNode struct {
    Type string
    Identifier string 
    Value interface{}
}

type BinOpNode struct {
    Type string
    Op string 
    Operand []interface{}
}

type IntNode struct {
    Type string
    Value int
}

type VarAcessNode struct {
    Type string 
    Identifier string
}

type StringNode struct {
    Type string 
    Value string
}

type EmptyNode struct {
    Type string
}
