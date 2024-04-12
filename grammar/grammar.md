# expression-lang grammar

## Grammar rules definition

1. declare variables


```
variable      -> STRING 

add_variable  -> variable, add_variable
                | variable
variables     -> "declare" add_variable

expression -> variables



