package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{l: lexer, errors: []string{}}

	// read two tokens to initialize both curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curError(t ...token.TokenType) {
	msg := fmt.Sprintf("expected current token to be one of %s, got %s instead", t, p.curToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram build a data structure (ast, abstract syntax tree) from the lexer output.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken() // skip semicolon
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		// error by returning nil
		p.peekError(token.IDENT)
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.peekError(token.ASSIGN)
		return nil
	}

	// TODO: implement expression parsing
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	if p.curToken.Type != token.INT && p.curToken.Type != token.IDENT {
		p.curError(token.INT, token.IDENT)
		return nil
	} else if p.curTokenIs(token.INT) {
		intValue, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			panic("Unexpected error converting integer string to int")
		}
		stmt.ReturnValue = &ast.IntExpression{Token: p.curToken, Value: intValue}
	} else if p.curTokenIs(token.IDENT) {
		if !p.peekTokenIs(token.LPAREN) {
			stmt.ReturnValue = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		} else {
			// TODO: Expression parsing
			for !p.curTokenIs(token.SEMICOLON) {
				p.nextToken()
			}
		}
	}
	return stmt
}
