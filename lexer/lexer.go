package lexer

import (
	"bufio"
	"os"
	"unicode"
)


type tokenType string 
const (
  //reserved 
  DECLARE="declare"
  EVALUATE="evaluate"

  //symbols
  COMMA="comma"
  PIPE="OR"
  E="AND"
  CIRCUMFLEX="XOR"
  EXCLAM="NOT"
 
  PAREN_LEFT="left parenthesis"
  PAREN_RIGHT="right parenthesis"

  SHIFT_LEFT="shift left"
  SHIFT_RIGHT="shift left"
  
  NUMBER="number"
  VARIABLE="variable"
)


type token struct{
  token string;
  token_type tokenType;
}



func Lexer(file *os.File) []token {
  var tokens []token = make([]token, 0);  
  scanner := bufio.NewScanner(file);

  for scanner.Scan(){
    line := scanner.Text()
    var actualChar int = 0


    for(actualChar <= len(line)-1){
      char := line[actualChar]

      switch(char){
        case ',': tokens = append(tokens, token{ token:",", token_type: COMMA })
        case '|': tokens = append(tokens, token{ token:"|", token_type: PIPE })
        case '&': tokens = append(tokens, token{ token:"&", token_type: E })
        case '^': tokens = append(tokens, token{ token:"^", token_type: CIRCUMFLEX })
        case '!': tokens = append(tokens, token{ token:"!", token_type: EXCLAM })
        case ')': tokens = append(tokens, token{ token:")", token_type: PAREN_RIGHT })
        case '(': tokens = append(tokens, token{ token:"(", token_type: PAREN_LEFT })
        case '>':
          next := getNext(actualChar, line);
          if next == ">" {
            tokens = append(tokens, token{ token:">>", token_type: SHIFT_RIGHT })
            actualChar += 2
            continue
          }
        case '<':
          next := getNext(actualChar, line);
          if next == "<" {
            tokens = append(tokens, token{ token:"<<", token_type: SHIFT_LEFT })
            actualChar += 2
            continue
          }

        case 'd':
          substring := line[actualChar:actualChar+7]
          if substring == "declare"{
            tokens = append(tokens, token{ token:"declare", token_type: DECLARE })
            actualChar += 8
            continue
          }
        
        case 'e':
          substring := line[actualChar:actualChar+8]
          if substring == "evaluate"{
            tokens = append(tokens, token{ token:"evaluate", token_type: EVALUATE })
            actualChar += 9 
            continue
          }
          

        default:
      
          if unicode.IsDigit(rune(char)){
            var number string;
            for(unicode.IsDigit(rune(char))){
              number += string(char)
              actualChar += 1
              
              if actualChar > len(line)-1{
                break
              }
              char = line[actualChar]
            }
            tokens = append(tokens, token{ token:number, token_type: NUMBER})
            continue 
          
          } else if unicode.IsLetter(rune(char)) || char == '_' {
            var variable string;
            for(unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)) || char == '_'){
              variable += string(char)
              actualChar += 1
              
              if actualChar > len(line)-1{
                break
              }
              char = line[actualChar]
            }

            tokens = append(tokens, token{ token:variable, token_type: VARIABLE})

            continue 

        } 


      }

      actualChar += 1
    }
  }
  return tokens;
}

func getNext(i int, line string) string {
  if i >= len(line)-1 {
    return "" 
  }
  
  return string(line[i+1])

}
