# Grammar 
prog = expr*

expr = atom | list

atom = number 
     = string
     = identifier 
     = - (expr)

list = (OP expr expr) ;OP = +/-/*
     = ( setf identifier expr)
     = ( identifier {expr}) ; function call 
     = ( defun identifier ({identifier}) expr)
     = ( if expr expr expr) ; condition then else
     = ( block expr* )
     = ( expr )


