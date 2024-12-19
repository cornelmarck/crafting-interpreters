package testdata

import "embed"

//go:embed *.lox
var TestCases embed.FS
