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
	return nil
}