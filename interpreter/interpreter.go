package interpreter 

import (
    "fmt"
    "strconv"
)

func handlePrint(i *Interpreter, funcNode FunctionCallNode) (interface{}, bool) {
    var output string

    for _,element := range funcNode.Parameters {
        v, err := eval(i,element)
        if err {return nil,true}

        switch val := v.(type) {
            case string:
                output += val
            case int:
                output += strconv.Itoa(val)
        } 
    }

    fmt.Println(output)
    return output, false
}

//var DefaultFunctions = map[string]func(FunctionCallNode) (interface{}, bool) { 
    //"print":handlePrint,
//}



type Interpreter struct {
    AST []interface{}
    Node interface{}
    Index int
    Variables map[string]interface{}
    Functions map[string]interface{}

}

func proceed(i *Interpreter){
    i.Index++
    i.Node = i.AST[i.Index]
}


func Interpret(i *Interpreter) bool {
    for i.Index < len(i.AST) - 1 {
        
        _,err := eval(i,i.Node)
        if err { return true } 
        proceed(i)
    }
    return false
}

func eval(i *Interpreter, n interface{}) (interface{}, bool) {
    switch node := n.(type) {
        case FunctionCallNode:
            value, err := evalFuncCall(i,node)
            if err {return nil, true}
            return value, false
        case VarAssignNode:
            value,err := evalVarAssign(i,node)
            if err {return nil, true}
            return value, false
        case BinOpNode:
            value,err := evalBinOp(i,node)
            if err {return nil, true}
            return value, false
        case VarAcessNode:
            return i.Variables[node.Identifier], false
        case IntNode:
            return node.Value, false
        case StringNode:
            return node.Value, false
    }
    return nil, true
}



func evalVarAssign(i *Interpreter, assignNode VarAssignNode) (interface{}, bool) {
    identifier := assignNode.Identifier
    value,err := eval(i,assignNode.Value)
    if err {return nil, true}
    i.Variables[identifier] = value
    return value, false

}


func evalFuncCall(i *Interpreter, funcNode FunctionCallNode) (interface{}, bool) {
    switch funcNode.Identifier {
        case "print":
            value, err := handlePrint(i,funcNode)
            if err {return nil, true}
            return value, false
    }

    return nil, false
}


func evalBinOp(i *Interpreter,opNode BinOpNode) (interface{}, bool) {
    initialOperand, err := eval(i,opNode.Operand[0])
    if err {return nil, true}

    switch initialOperand.(type) {
        case string:
            var stringResult string = initialOperand.(string)

            for _,op := range opNode.Operand[1:] {
                evalulatedOp, err := eval(i,op)
                if err {return nil, true}

                stringResult += evalulatedOp.(string)
            }

            return stringResult, false
    
        case int:
            var intResult int = initialOperand.(int)

            for _,op := range opNode.Operand[1:] {
                evalulatedOp, err := eval(i,op)
                if err {return nil, true}

                intResult = calculate(opNode.Op, intResult, evalulatedOp).(int)
            }

            return intResult, false
    }


    return nil, true
}

func calculate(op string, result interface{}, value interface{}) interface{}{
    switch op {
        case TT_ADD:
            return result.(int) + value.(int)
        case TT_SUB:
            return result.(int) - value.(int)
        case TT_MUL:
            return result.(int) * value.(int)
        case TT_DIV:
            return result.(int) / value.(int)
    }

    return nil
}
