package interpreter 

import (
    "fmt"
    "strconv"
    "math/rand"
    "time"
)


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
        
        _,err := eval(i,i.Node, i.Variables)
        if err { return true } 
        proceed(i)
    }
    return false
}

func eval(i *Interpreter, n interface{}, scope map[string]interface{}) (interface{}, bool) {
    switch node := n.(type) {
        case IfNode:
            value, err := evalIf(i,node, scope)
            if err {return nil,true}
            return value,false
        case WhileNode:
            value, err := evalWhile(i,node, scope)
            if err {return nil, true}
            return value, false
        case FunctionDefenitionNode:
            i.Functions[node.Identifier] = node
            return nil, false
        case FunctionCallNode:
            value, err := evalFuncCall(i,node, scope)
            if err {return nil, true}
            return value, false
        case VarAssignNode:
            value,err := evalVarAssign(i,node, scope)
            if err {return nil, true}
            return value, false
        case BinOpNode:
            value,err := evalBinOp(i,node, scope)
            if err {return nil, true}
            return value, false
        case BlockNode:
            value,err := evalBlock(i,node, scope)
            if err {return nil, true}
            return value,false
        case VarAcessNode:
            return scope[node.Identifier], false
        case IntNode:
            return node.Value, false
        case StringNode:
            return node.Value, false
    }
    return nil, true
}



func evalWhile(i *Interpreter, whileNode WhileNode, scope map[string]interface{}) (interface{}, bool) {
    var condition interface{}
    var value interface{}
    var err bool = false

    condition,err = eval(i, whileNode.Condition, scope)
    if err {return nil, true}

    for condition.(int) != 0 {
        value,err = eval(i,whileNode.Consequence, scope)
        if err {return nil, true}

        condition, err = eval(i,whileNode.Condition, scope)
        if err {return nil, true}
    }

    return value, false

}


func evalBlock(i *Interpreter, blockNode BlockNode, scope map[string]interface{}) (interface{}, bool) {
    var err bool = false
    var value interface{}

    for _,node := range blockNode.Block {
        value, err = eval(i,node,scope)
        if err {return nil, true}
    }

    return value, false
}


func evalIf(i *Interpreter, ifNode IfNode, scope map[string]interface{}) (interface{}, bool) {
    condition,err := eval(i, ifNode.Condition, scope)
    if err {return nil, true}

    if condition == 0 {
        value,err := eval(i,ifNode.Alternative,scope)
        if err {return nil, true}
        return value, false
    } else {
        value, err := eval(i,ifNode.Consequence,scope)
        if err {return nil, true}
        return value, false
    }

    return nil, true
}


func evalVarAssign(i *Interpreter, assignNode VarAssignNode, scope map[string]interface{}) (interface{}, bool) {
    identifier := assignNode.Identifier
    value,err := eval(i,assignNode.Value,scope)
    if err {return nil, true}
    scope[identifier] = value
    return value, false

}


func evalFuncCall(i *Interpreter, funcNode FunctionCallNode, scope map[string]interface{}) (interface{}, bool) {
    switch funcNode.Identifier {
        case "print":
            value, err := handlePrint(i,funcNode,scope)
            if err {return nil, true}
            return value, false
        case "len":
            value,err := handleLen(i,funcNode,scope)
            if err {return nil, true}
            return value, false
        case "rnd":
            value,err := handleRnd(i,funcNode,scope)
            if err {return nil, true}
            return value, false
    }

    if function,contained := i.Functions[funcNode.Identifier]; contained {
        var localScope = make(map[string]interface{})
        parameterNames := function.(FunctionDefenitionNode).Parameters
        parameterValues := funcNode.Parameters

        for index,name := range parameterNames {
            value := parameterValues[index]
            evalValue, err := eval(i,value,i.Variables)
            if err {return nil, true}
            localScope[name.(ParameterNode).Identifier] = evalValue
        }

        value, err := eval(i,function.(FunctionDefenitionNode).Block,localScope)
        if err {return nil, true}

        return value,false
    }

    fmt.Println("Function:",funcNode.Identifier,"does not exist")
    return nil, true
}


func evalBinOp(i *Interpreter,opNode BinOpNode, scope map[string]interface{}) (interface{}, bool) {
    initialOperand, err := eval(i,opNode.Operand[0],scope)
    if err {return nil, true}

    var COMPARISONS = []string{TT_EE,TT_NE,TT_GT,TT_LT,TT_GTE,TT_LTE,}
    
    if contains(COMPARISONS, opNode.Op) {
        operand1, err := eval(i,opNode.Operand[0],scope)
        operand2, err  := eval(i,opNode.Operand[1],scope)
        if err {return nil, true}

        var result int = 0

        switch operand1.(type) {
            case int:
                switch opNode.Op {
                    case TT_EE:
                        if operand1.(int) == operand2.(int) {result = 1}
                    case TT_NE:
                        if operand1.(int) != operand2.(int) {result = 1}
                    case TT_LT:
                        if operand1.(int) < operand2.(int) {result = 1}
                    case TT_GT:
                        if operand1.(int) > operand2.(int) {result = 1}
                    case TT_LTE:
                        if operand1.(int) <= operand2.(int) {result = 1}
                    case TT_GTE:
                        if operand1.(int) >= operand2.(int) {result = 1}
                }

            case string:
                switch opNode.Op {
                    case TT_EE:
                        if operand1.(string) == operand2.(string) {result = 1}
                    case TT_NE:
                        if operand1.(string) != operand2.(string) {result = 1}
                }

        }

        return result, false

    } 

    switch initialOperand.(type) {
        case string:
            var stringResult string = initialOperand.(string)

            for _,op := range opNode.Operand[1:] {
                evalulatedOp, err := eval(i,op,scope)
                if err {return nil, true}

                stringResult += evalulatedOp.(string)
            }

            return stringResult, false
    
        case int:
            var intResult int = initialOperand.(int)

            for _,op := range opNode.Operand[1:] {
                evalulatedOp, err := eval(i,op,scope)
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

func contains(array []string, element string) bool {
    for _,e := range array {
        if e == element {
            return true
        }
    }

    return false
}


// DEFAULT FUNCTIONS 

func handlePrint(i *Interpreter, funcNode FunctionCallNode, scope map[string]interface{}) (interface{}, bool) {
    var output string

    for _,element := range funcNode.Parameters {
        v, err := eval(i,element,scope)
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

func handleLen(i *Interpreter, funcNode FunctionCallNode, scope map[string]interface{}) (interface{}, bool) {
    rawString, err := eval(i, funcNode.Parameters[0],scope)
    if err {return nil, true}
    return len(rawString.(string)), false
}

func handleRnd(i *Interpreter, funcNode FunctionCallNode, scope map[string]interface{}) (interface{}, bool) {
    rand.Seed(time.Now().UnixNano())
    parameters := funcNode.Parameters

    min := 0
    max := 0
    if (len(parameters) >= 2) {
        minVal,err := eval(i, parameters[0], scope)
        maxVal,err := eval(i, parameters[1], scope)
        if err {return nil, true}
        min = minVal.(int)
        max = maxVal.(int)

    } else {
        maxVal, err := eval(i, parameters[0], scope)
        if err {return nil, true}
        max = maxVal.(int)
    }

    return rand.Intn(max - min) + min, false

}
