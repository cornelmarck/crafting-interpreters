# Crafting Interpreters

 Crafting Interpreters (2021) is a great book by written by Robert Nystrom. It is availabe for free at [craftinginterpreters.com](https://craftinginterpreters.com/contents.html).

This repo contains my implementation of the Lox language and interpreter. I worked on this repo while studying the book, re-implementing the Java codebase in idiomatic Go.

## Grammar

Lox is described by the following context-free grammar.
```
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expression ")" ;
```

## Data types

Lox has the following data types:
- Booleans
- Numbers: Both integers and decimals are implemented as 64 bit floats.
- Strings
- Nil
