# Lisp-Interpreter
> A languge based on LISP syntax, NOT strictly lisp, I changed some things that in my opinion make the language more useable, this was a quick project to learn the basics of golang
## Overview of Syntax
*every block returns something, e.g. if will return either the  consequence or the alternative, print will return the string, variable assignments will return the value that is being assigned.
### Variable assignment
(setf *variable name* *value)
  - e.g. `(setf x 10)`
  - `(setf result (+ 10 x))
  
### Function Calls
(*identifier* *parameters*)
  - e.g. `(print "Your name is " name " and you are " age " years old")`
  - e.g. `(input "Enter name: ")
 
### If Statements
*consequence and alternative can only be 1 expression, so a block statement can be used to perform multiple*
(if *condition* *consequence* *alternative*)
  - e.g. `(print (if (== answer 10) "Correct" "Wrong")`
  - e.g. `(if (== answer 10) (block (print "Correct") (setf correct 1)) (block (print "Wrong") (setf correct 0))`

### Block statements
*they contain multiple statements and return the result of the last
(block expr*)
  - e.g. `(block (print "That answer was correct") (setf correct 1))`
     - This would return 1
     
 ### While Statements
 (while *condition* *consequence*)
  - e.g. `(while (< i 10) (block (setf i (+ i 1)) (print i))`
  
 ### Default functions
 - print  - print
 - input - input
 - intin - int input 
 - rnd - random number 
- len - length of string

*Example programs in examples folder*
