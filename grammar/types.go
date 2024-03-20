package grammar

import (
  "exp/exp/lexer"
)

type Expression struct{
  exp any //TODO: improve that
}

type Literal struct{
  literal lexer.Token 
}

type Unary struct{
  op lexer.Token 
  exp Expression
} 

type Binary struct {
  left_exp Expression
  op Operator
  right_exp Expression
}

type Operator struct{
  op lexer.Token 
}

type Grouping struct{
  left_gr lexer.Token
  exp Expression
  right_gr lexer.Token
}
