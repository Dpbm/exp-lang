package lexer

import (
	"bufio"
	"os"
	"unicode"
)

func Lexer(file *os.File) *[]Token {
  var tokens []Token = make([]Token, 0);  
  scanner := bufio.NewScanner(file);

  for scanner.Scan(){
    line := scanner.Text()
    var actualChar int = 0


    for(actualChar <= len(line)-1){
      char := line[actualChar]

      switch(char){
        case ',': tokens = append(tokens, Token{ Symbol:",", Type: COMMA })
        case '|': tokens = append(tokens, Token{ Symbol:"|", Type: PIPE })
        case '&': tokens = append(tokens, Token{ Symbol:"&", Type: E })
        case '^': tokens = append(tokens, Token{ Symbol:"^", Type: CIRCUMFLEX })
        case ')': tokens = append(tokens, Token{ Symbol:")", Type: PAREN_RIGHT })
        case '(': tokens = append(tokens, Token{ Symbol:"(", Type: PAREN_LEFT })
        case '-': tokens = append(tokens, Token{ Symbol:"-", Type: MINUS })
        
        case '>':
          next := getNext(actualChar, line);
          if next == ">" {
            tokens = append(tokens, Token{ Symbol:">>", Type: SHIFT_RIGHT })
            actualChar += 2
            continue
          }
        
        case '<':
          next := getNext(actualChar, line);
          if next == "<" {
            tokens = append(tokens, Token{ Symbol:"<<", Type: SHIFT_LEFT })
            actualChar += 2
            continue
          }
        
        case '=':
          next := getNext(actualChar, line);
          if next == "=" {
            tokens = append(tokens, Token{ Symbol:"==", Type: EQUAL })
            actualChar += 2
            continue
          }else{
            tokens = append(tokens, Token{ Symbol:"=", Type: ASSIGN })
          }
        
        case '!':
          next := getNext(actualChar, line);
          if next == "=" {
            tokens = append(tokens, Token{ Symbol:"!=", Type: DIFFERENT })
            actualChar += 2
            continue
          }else{
            tokens = append(tokens, Token{ Symbol:"!", Type: EXCLAM })
          }
        
        case 'T':
          substring := line[actualChar:actualChar+4]
          if substring == "TRUE"{
            tokens = append(tokens, Token{ Symbol:"TRUE", Type: TRUE })
            actualChar += 5
            continue
          }
        
        case 'F':
          substring := line[actualChar:actualChar+5]
          if substring == "FALSE"{
            tokens = append(tokens, Token{ Symbol:"FALSE", Type: FALSE })
            actualChar += 6
            continue
          }

        case 'd':
          substring := line[actualChar:actualChar+7]
          if substring == "declare"{
            tokens = append(tokens, Token{ Symbol:"declare", Type: DECLARE })
            actualChar += 8
            continue
          }
        
        case 'e':
          substring := line[actualChar:actualChar+8]
          if substring == "evaluate"{
            tokens = append(tokens, Token{ Symbol:"evaluate", Type: EVALUATE })
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
            tokens = append(tokens, Token{ Symbol:number, Type: NUMBER})
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

            tokens = append(tokens, Token{ Symbol:variable, Type: VARIABLE})
            continue 
        } 
      }

      actualChar += 1
    }

    tokens = append(tokens, Token{ Symbol:"", Type: END})

  }


  return &tokens;
}

func getNext(i int, line string) string {
  if i >= len(line)-1 {
    return "" 
  }
  
  return string(line[i+1])

}
