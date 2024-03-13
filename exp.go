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
  data, err := os.ReadFile(file)
  
  if err != nil {
    fmt.Println("Invalid file!")
    os.Exit(1)
  }

  lexer.Lexer(string(data));  

}
