# Lexer
(defun compare x (
	if (> x 10) "Large" "Small"
))

*Format to single line*
while currentTok != null:
	if currentTok == "STRING": return parseString()	
	else if currentTok == "INT": return parseInt()
	else if currentTok == "LETTER": return parseIdentifier()
	else if currentTok == "(": return LPAREN
	else if currentTok == ")": return TPAREN
	tok.advance()
	
**Format**
(defun compare x (if (> x 10) "Large" "Small"))

**Tokenise**
LPAREN KEYWORD ID LPAREN KEYWORD LPAREN GT ID INT RPAREN STRING STRING RPAREN RPAREN 