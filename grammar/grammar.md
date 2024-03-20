# EXP-lang grammar

## Definitions

```
literals: Numbers, Strings(variables), Booleans
unary: !, -, etc.
binary: ==, !=, &, >>, <<, etc.
grouping: (, )
```

---

## Symbols table

```
SYMBOLS {
  literal: "TRUE", "FALSE", NUMBERS, VARIABLES (STRING)

  unary: "-", "!", "declare", "evaluate"
  
  binary: "|", "&", "^", ">>", "<<", "=", "==", "!=", ","  

  grouping: "(", ")"
}
```

---

## Grammar rules definition

```
expression  -> literal
              | unary
              | binary
              | grouping;
literal     ->  NUMBER | STRING | "TRUE" | "FALSE";
grouping    -> "(" expression ")";
unary       -> ( "!" | "-" | "declare" | "evaluate" ) expression;
binary      -> expression operator expression;
operator    -> "|" | "&" | "^" | ">>" | "<<" | "=" | "==" | "!=" | ",";
```
