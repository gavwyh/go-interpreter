package ast

import "github.com/gavwyh/go-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

// root node of every AST
type Program struct {
	Statements []Statement
}

// Implements expression as identifiers can produce values -> anyth produces val = expression
type Identifier struct {
	Token token.Token
	Value string
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }