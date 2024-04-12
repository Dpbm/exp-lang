package parser

import (
	"errors"
	"exp/exp/lexer"
	"fmt"
	"os"
)

func Parse(tokens *[]lexer.Token) Tree {
  var i uint64  = 0 
  return expression(tokens, &i);
}

func expression(tokens *[]lexer.Token, i *uint64) Tree{
  var tree Tree = Tree{}
  token := next(tokens, i)

  //variable
  if(token.Type == lexer.DECLARE){
    tree.Value = token
    walk(i)
    value, err := add_variable(tokens, i)
    if err != nil{
      fmt.Println(err);
      os.Exit(1); 
    }

    tree.Right = value
  }
  return tree
}

func add_variable(tokens *[]lexer.Token, i *uint64) (*Node,error){
  token := next(tokens, i)
  previous_token := previous(tokens, i)
 
  switch token.Type {
    case lexer.COMMA:
      if(previous_token.Type != lexer.VARIABLE){
        return nil, errors.New(fmt.Sprintf("invalid ',' after '%s'", previous_token.Symbol))
      }

    case lexer.VARIABLE:
      if(previous_token.Type != lexer.COMMA && previous_token.Type != lexer.DECLARE){
        return nil, errors.New(fmt.Sprintf("invalid '%s' after '%s'", token.Symbol, previous_token.Symbol))
      }
  
    case lexer.END:
      if(previous_token.Type == lexer.DECLARE){
        return nil, errors.New("invalid EOF after declare")
      }

    default:
      return nil, errors.New(fmt.Sprintf("invalid token '%s'", token.Symbol))
    
  }
  
  var node Node = Node{}
  
  if(token.Type == lexer.END){
    node.Value = token
    return &node, nil
  }

  node.Value = token
  walk(i)
  value, err := add_variable(tokens, i)
  if err == nil{
    node.Right = value
    return &node, nil
  }else{
    return nil, err
  }
}


func next(tokens *[]lexer.Token, i *uint64) lexer.Token{
  return (*tokens)[*i]
}

func previous(tokens *[]lexer.Token, i *uint64) lexer.Token{
  return (*tokens)[(*i)-1]
}

func walk(i *uint64){
  *i++
}


