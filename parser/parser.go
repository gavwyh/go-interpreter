package parser

import (
	"fmt"
	"strconv"

	"github.com/gavwyh/go-interpreter/token"
	"github.com/gavwyh/go-interpreter/lexer"
	"github.com/gavwyh/go-interpreter/ast"
)

const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT // *
	PREFIX // -X or !X
	CALL // myFunction(X)
)
	

type Parser struct {
	lexer *lexer.Lexer
	errors []string

	curToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)
	

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer : lexer,
		errors: []string{},
	}

	// Read two tokens to set both curToken & peekToken -> required to determine if 
		// at EOL or start of arithmetic expression e.g 5; vs 5 * 5;
	parser.nextToken()
	parser.nextToken()

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(token.INT, parser.parseIntegerLiteral)

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
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
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

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: parser.curToken}

	statement.Expression = parser.parseExpression(LOWEST)

	if parser.isPeekToken(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExpression := prefix()

	return leftExpression
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{Token: parser.curToken}

	value, err := strconv.ParseInt(parser.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as an integer", parser.curToken.Literal)
		parser.errors = append(parser.errors, msg);
		return nil;
	}

	literal.Value = value;
	return literal;
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

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn;
}

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn;
}