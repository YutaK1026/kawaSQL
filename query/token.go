package query

import "strings"

type TokenKind int

type Tokenizer struct {
	input string
	pos   int
}

func (t TokenKind) String() string {
	return tokens[t]
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input: input,
		pos:   0,
	}
}