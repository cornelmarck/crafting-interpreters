package main

import "errors"

var (
	ErrInvalidOperand = errors.New("invalid operand")
	ErrDivideByZero   = errors.New("divide by zero")
)
