package ast

import (
	"bytes"
	"monkey/token"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is a root node of ASTs
// Program consists of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// let x;
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	// return;
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String() + ";"
	}
	return ""
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) TokenLiteral() string {
	return strconv.FormatInt(i.Value, 10)
}
func (i *IntegerLiteral) String() string {
	return i.TokenLiteral()
}

type PrefixOperator struct {
	Token token.Token
	Value Expression
}

func (po *PrefixOperator) expressionNode() {}
func (po *PrefixOperator) TokenLiteral() string {
	return po.Token.Literal
}
func (po *PrefixOperator) String() string {
	var out bytes.Buffer
	out.WriteString("(" + " ")
	out.WriteString(po.Token.Literal)
	out.WriteString(po.Value.String())
	out.WriteString(" " + ")")
	return out.String()
}

type InfixOperator struct {
	Token    token.Token // Operator token
	Left     Expression
	Operator string
	Right    Expression
}

func (io *InfixOperator) expressionNode() {}
func (io *InfixOperator) TokenLiteral() string {
	return io.Token.Literal
}
func (io *InfixOperator) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(io.Left.String())
	out.WriteString(" " + io.Operator + " ")
	out.WriteString(io.Right.String())
	out.WriteString(")")
	return out.String()
}
