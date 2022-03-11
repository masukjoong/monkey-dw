package ast

import (
	"monkey/token"
	"testing"
)

func TestLetStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	assertEquals(t, program, "let myVar = anotherVar;")
}

func TestLetStatementWithNilValue_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: nil,
			},
		},
	}

	assertEquals(t, program, "let myVar = ;")
}

func TestReturnStatement_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token:       token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
			},
		},
	}

	assertEquals(t, program, "return x;")
}

func assertEquals(t *testing.T, program *Program, expected string) {
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestPrefixOperator_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Token: token.Token{Type: token.MINUS, Literal: "-"},
				Expression: &PrefixOperator{Token: token.Token{Type: token.MINUS, Literal: "-"}, Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "123"}, Value: "123"},
				},
			},
		},
	}

	assertEquals(t, program, "-123;")
}
