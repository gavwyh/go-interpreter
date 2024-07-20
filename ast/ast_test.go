package ast

import (
	"testing"

	"github.com/gavwyh/go-interpreter/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "y"},
					Value: "y",
				},
			},
		},
	}

	if program.String() != "let x = y;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}