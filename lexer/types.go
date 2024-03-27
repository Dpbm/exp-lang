package lexer

type Token struct{
  Symbol string
  Type LexicalTokenType 
}

type LexicalTokenType string 
const (
  //reserved 
  DECLARE="declare"
  EVALUATE="evaluate"
  TRUE="true"
  FALSE="false"

  //symbols
  COMMA="comma"
  PIPE="OR"
  E="AND"
  CIRCUMFLEX="XOR"
  EXCLAM="NOT"
  EQUAL="equal sign"
  ASSIGN="assign"
  DIFFERENT="different sign"
  MINUS="minus"
 
  PAREN_LEFT="left parenthesis"
  PAREN_RIGHT="right parenthesis"

  SHIFT_LEFT="shift left"
  SHIFT_RIGHT="shift right"
  
  NUMBER="number"
  VARIABLE="variable"
)

