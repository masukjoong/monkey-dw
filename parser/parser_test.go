package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if stmt.TokenLiteral() != "let" {
			t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		}
		letStmt, ok := (stmt).(*ast.LetStatement)
		if !ok {
			t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		}
		if letStmt.Name.Value != tt.expectedIdentifier {
			t.Errorf("letStmt.Name.ReturnValue not %s. got=%s", tt.expectedIdentifier, letStmt.Name.Value)
		}
		if letStmt.Name.TokenLiteral() != tt.expectedIdentifier {
			t.Errorf("letStmt.TokenLiteral not %s. got=%s", tt.expectedIdentifier, letStmt.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	code := `
return 10;
return 20;
return 993322;
`
	l := lexer.New(code)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	tests := []struct {
		expectedExpression string
	}{
		{"10"},
		{"20"},
		{"993322"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if stmt.TokenLiteral() != "return" {
			t.Errorf("s.TokenLiteral not 'return'. got=%q", stmt.TokenLiteral())
		}

		rs, ok := (stmt).(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not *ast.ReturnStatement. got=%T", stmt)
		}

		if rs.ReturnValue.TokenLiteral() != tt.expectedExpression {
			t.Errorf("rs.ReturnValue.TokenLiteral() not %s. got=%s", tt.expectedExpression, rs.ReturnValue.TokenLiteral())
		}
	}
}