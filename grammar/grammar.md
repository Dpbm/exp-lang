# expression-lang grammar

## Grammar rules definition

```
variable    -> STRING; 
values -> NUMBER | "TRUE" | "FALSE"
base   ->  values | variable | "(" expression ")";
comparation   -> "==" | "!=";
binary   -> "|" | "&" | "^" | ">>" | "<<";

declaration   -> "declare" declaration 
            | variable "," declaration
            | variable;

attribution   -> variable "=" base;
evaluation  -> "evaluate" expression;

modification   -> ( "!" | "-" ) modification 
          | base;

expression   ->  modification (comparation | binary) expression
           | modification; 
```
