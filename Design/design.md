# Design
## Variables
this.interpretToken( scope = global / local)

### Declare

VarDeclare = {
    srting identifier 
    value = node
    scope = scope ; if being parsed in main then global else function
}

### Access 
VarAccess = {
    string identifier
    scope = scope
    eval(interpreter)= {
            if scope == global { 
                return interpreter.global[identifier]    
            }else{
                return interpreter.local[identifier]
            }
        }
}

### Scope
*scope will be parsed into interpretToken*

global = []
local = []

## Functions

### Function Call
FunctionNode = {
    string Identifier 
    variables = (local + global)
    parameters = loop&assign(parameters) ; assign parameters to local
}

### Declare 
FunctionNode = {
        string Identifier 
        bodyNody body 
        array paramters
    }

## Examples
### Comparison
(defun compare (x) (
    if(> x 10) (
        if (> x 15) "Huge" "Large"
    ) 
    "Small"
))

(setf num 100)
(print (compare x))

global = {num:100}
local = {x:100}

functionCall = compare.eval() = if.eval() = if condition.eval() then return "Huge" else "Large"
functionDec = functions.append(name, parameters, body)


block(
    func(
		parameters = [x]
        consequence = block(
            if(
                condition = x > 10
                consequence = block(
                    if (
                        if(
                            condition = x > 15
                            consequence = "Huge"
                            alternative = "Large"
                        )
                    )
                )
                alternative = "Large"
            )
        )
    )
)


### Addition
(+ (- 10 2) 15)

block (
    expr (
        ADD (
            MINUS(
                10 - 2
            ),
            NUM(
            	15
            )
        )  
   	)
)


### Print
(print (+ "Hello" "World"))

block(
    expr(
        func(
            ADD(
                STRING("Hello")
                STING("World")
            )
        )
    )
