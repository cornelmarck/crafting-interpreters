# Crafting Interpreters

 Crafting Interpreters is a book on programming language design by Robert Nystrom. It is availabe for free at [craftinginterpreters.com](https://craftinginterpreters.com/contents.html).

This repo contains my implementation of the Lox language and interpreter. I worked on this repo while studying the book, re-implementing the Java interpreter in Go, and the bytecode VM in Zig instead of C.


## The Lox language

Lox is an interpreted language that was designed within the context of the book. It is a dynamic language that supports classes, functions, variables, statements and expressions.

Lox is described by a context-free grammer [(link here)](https://craftinginterpreters.com/appendix-i.html).


Lox has the following data types:
- Booleans
- Numbers: Both integers and decimals are implemented as 64 bit floats.
- Strings
- Nil
