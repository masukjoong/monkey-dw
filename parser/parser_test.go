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

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := (stmt).(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("s.TokenLiteral not 'return'. got=%q", stmt.TokenLiteral())
		}
	}
}

func TestParser_parseExpressionStatement(t *testing.T) {
	code := `
1;
a;
let a = 10;
return 10;
`
	l := lexer.New(code)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	// print statements for debugging purposes
	for _, stmt := range program.Statements {
		println(stmt.String())
	}

	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d", len(program.Statements))
	}
}

func TestParser_ParsePrefixOperator(t *testing.T) {
	code := `
-123;
!test;
`
	l := lexer.New(code)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	answers := []string{"( -123 )", "( !test )"}

	for i, stmt := range program.Statements {
		println(stmt.String())

		es, ok := (stmt).(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("statement is not an ast.ExpressionStatement")
		}

		po, ok := es.Expression.(*ast.PrefixOperator)
		if !ok {
			t.Errorf("Expression statement's expression is not PrefixOperator")
		}

		if po.String() != answers[i] {
			t.Errorf("got = %q, expected = %q", po.String(), answers[i])
		}
	}
}

func TestParser_ParseInfixOperator(t *testing.T) {
	code := `
3 + 4;
-3 + 4;
`
	l := lexer.New(code)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	answers := []string{"(3 + 4);", "((-3)+4);"}

	for i, stmt := range program.Statements {
		println(i, stmt.String())

		if stmt.String() != answers[i] {
			t.Errorf("got = %q, expected = %q", stmt.String(), answers[i])
		}
	}
}
