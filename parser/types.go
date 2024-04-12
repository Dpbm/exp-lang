package parser

import (
	"exp/exp/lexer"
)

type Node struct{
  Value  lexer.Token
  Left *Node
  Right *Node
}

type Tree = Node
