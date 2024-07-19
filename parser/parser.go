package parser

import (
	"fmt"

	"github.com/gavwyh/go-interpreter/token"
	"github.com/gavwyh/go-interpreter/lexer"
	"github.com/gavwyh/go-interpreter/ast"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken token.Token
	peekToken token.Token

	errors []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer : lexer,
		errors: []string{},
	}

	// Read two tokens to set both curToken & peekToken -> required to determine if 
		// at EOL or start of arithmetic expression e.g 5; vs 5 * 5;
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) addError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got=%s",
			t, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
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
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: parser.curToken}

	if !parser.peekExpected(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	if !parser.peekExpected(token.ASSIGN) {
		return nil;
	}

	for !parser.isCurToken(token.SEMICOLON) {
		parser.nextToken()
	}
	return statement;
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: parser.curToken}

	parser.nextToken()

	for !parser.isCurToken(token.SEMICOLON) {
		parser.nextToken()
	}
	return statement
}

func (parser *Parser) isCurToken(t token.TokenType) bool { return parser.curToken.Type == t }

func (parser *Parser) isPeekToken(t token.TokenType) bool { return parser.peekToken.Type == t }

func (parser *Parser) peekExpected(expectedType token.TokenType) bool {
	if parser.isPeekToken(expectedType) {
		parser.nextToken()
		return true
	}
	parser.addError(expectedType)
	return false;
}