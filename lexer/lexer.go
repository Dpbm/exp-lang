package lexer

import (
	"bufio"
	"os"
	"unicode"
)

func Lexer(file *os.File) []Token {
  var tokens []Token = make([]Token, 0);  
  scanner := bufio.NewScanner(file);

  for scanner.Scan(){
    line := scanner.Text()
    var actualChar int = 0


    for(actualChar <= len(line)-1){
      char := line[actualChar]

      switch(char){
        case ',': tokens = append(tokens, Token{ symbol:",", lexical_type: COMMA })
        case '|': tokens = append(tokens, Token{ symbol:"|", lexical_type: PIPE })
        case '&': tokens = append(tokens, Token{ symbol:"&", lexical_type: E })
        case '^': tokens = append(tokens, Token{ symbol:"^", lexical_type: CIRCUMFLEX })
        case ')': tokens = append(tokens, Token{ symbol:")", lexical_type: PAREN_RIGHT })
        case '(': tokens = append(tokens, Token{ symbol:"(", lexical_type: PAREN_LEFT })
        case '-': tokens = append(tokens, Token{ symbol:"-", lexical_type: MINUS })
        
        case '>':
          next := getNext(actualChar, line);
          if next == ">" {
            tokens = append(tokens, Token{ symbol:">>", lexical_type: SHIFT_RIGHT })
            actualChar += 2
            continue
          }
        
        case '<':
          next := getNext(actualChar, line);
          if next == "<" {
            tokens = append(tokens, Token{ symbol:"<<", lexical_type: SHIFT_LEFT })
            actualChar += 2
            continue
          }
        
        case '=':
          next := getNext(actualChar, line);
          if next == "=" {
            tokens = append(tokens, Token{ symbol:"<<", lexical_type: EQUAL })
            actualChar += 2
            continue
          }else{
            tokens = append(tokens, Token{ symbol:"=", lexical_type: ASSIGN })
          }
        
        case '!':
          next := getNext(actualChar, line);
          if next == "=" {
            tokens = append(tokens, Token{ symbol:"!=", lexical_type: DIFFERENT })
            actualChar += 2
            continue
          }else{
            tokens = append(tokens, Token{ symbol:"!", lexical_type: EXCLAM })
          }
        
        case 'T':
          substring := line[actualChar:actualChar+4]
          if substring == "TRUE"{
            tokens = append(tokens, Token{ symbol:"TRUE", lexical_type: TRUE })
            actualChar += 5
            continue
          }
        
        case 'F':
          substring := line[actualChar:actualChar+5]
          if substring == "FALSE"{
            tokens = append(tokens, Token{ symbol:"FALSE", lexical_type: FALSE })
            actualChar += 6
            continue
          }

        case 'd':
          substring := line[actualChar:actualChar+7]
          if substring == "declare"{
            tokens = append(tokens, Token{ symbol:"declare", lexical_type: DECLARE })
            actualChar += 8
            continue
          }
        
        case 'e':
          substring := line[actualChar:actualChar+8]
          if substring == "evaluate"{
            tokens = append(tokens, Token{ symbol:"evaluate", lexical_type: EVALUATE })
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
            tokens = append(tokens, Token{ symbol:number, lexical_type: NUMBER})
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

            tokens = append(tokens, Token{ symbol:variable, lexical_type: VARIABLE})

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
