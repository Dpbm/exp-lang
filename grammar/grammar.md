# EXP-lang grammar

## Grammar rules definition

```
var    -> STRING; 
values -> NUMBER | "TRUE" | "FALSE"
base   ->  values | var | "(" exp ")";
cmp   -> "==" | "!=";
bin   -> "|" | "&" | "^" | ">>" | "<<";

dec   -> "declare" dec 
            | var "," dec
            | var;

att   -> var "=" base;
eval  -> "evaluate" exp;

mod   -> ( "!" | "-" ) mod 
          | base;

exp   ->  mod (cmp | bin) exp
           | mod; 
```
