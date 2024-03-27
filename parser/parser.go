package parser

import (
	"errors"
	"exp/exp/grammar"
	"exp/exp/lexer"
	"fmt"
	"strconv"
)

func Parse(tokens *[]lexer.Token){
  for _, token := range *tokens{
    fmt.Println(token.Symbol);
  }
}

func values(token lexer.Token) (grammar.Literal, error) {
  switch token.Type{
    case lexer.TRUE:
      return true_token()
    case lexer.FALSE:
      return false_token()
    case lexer.NUMBER:
      return number_token(token)
    default:
      return grammar.Literal{}, errors.New("Invalid values token")
  }
}


func variable(token lexer.Token) (grammar.Literal, error) {
  if token.Type == lexer.VARIABLE{
    return variable_token(token)
  }

  return grammar.Literal{}, errors.New("Invalid variable")
}

func base(token lexer.Token) (grammar.Literal, error) {
  switch(token.Type){
    case lexer.TRUE:
      return true_token()
    case lexer.FALSE:
      return false_token()
    case lexer.NUMBER:
      return number_token(token)
    case lexer.VARIABLE:
      return variable_token(token)
    case lexer.PAREN_LEFT:
      //TODO
    default:
      return grammar.Literal{}, errors.New("Invalid base token")
  }
}

func true_token() (grammar.Literal,error){
  return grammar.Literal{Value: true},nil
}

func false_token() (grammar.Literal,error){
  return grammar.Literal{Value: false},nil
}

func variable_token(token lexer.Token) (grammar.Literal,error){
  return grammar.Literal{Value: token.Symbol},nil
}

func number_token(token lexer.Token) (grammar.Literal, error){
  value, err := strconv.Atoi(token.Symbol)
  if err == nil{
    return grammar.Literal{Value:value}, nil
  }else{
    return grammar.Literal{}, errors.New("Error on converting token to int")
  }
}
