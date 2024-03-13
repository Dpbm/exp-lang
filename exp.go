package main 

import (
  "fmt"
  "os"
  "exp/exp/lexer"
)

func main(){
  
  if len(os.Args) != 2 {
    fmt.Println("Usage: exp source_code_file.exp")
    os.Exit(1)
  }


  var file string = os.Args[1]
  file_o, err := os.Open(file)
  
  if err != nil {
    fmt.Println("Invalid file!")
    os.Exit(1)
  }

  tokens := lexer.Lexer(file_o);   
  defer file_o.Close()

  for _, token := range tokens{
    fmt.Println(token);
  } 
}
