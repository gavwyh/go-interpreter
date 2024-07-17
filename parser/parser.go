package parser

import (
	"github.com/gavwyh/go-interpreter/token"
	"github.com/gavwyh/go-interpreter/lexer"
	"github.com/gavwyh/go-interpreter/ast"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer : lexer}

	// Read two tokens to set both curToken & peekToken -> required to determine if 
		// at EOL or start of arithmetic expression e.g 5; vs 5 * 5;
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}
	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: parser.curToken}

	if !parser.nextExpected(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	if !parser.nextExpected(token.ASSIGN) {
		return nil;
	}

	for parser.curToken.Type != token.SEMICOLON {
		parser.nextToken()
	}
	return statement;
}

func (parser *Parser) nextExpected(expectedType token.TokenType) bool {
	if parser.peekToken.Type == expectedType {
		parser.nextToken()
		return true
	}
	return false;
} 