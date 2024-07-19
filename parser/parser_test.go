package parser

import (
	"testing"
	
	"github.com/gavwyh/go-interpreter/lexer"
	"github.com/gavwyh/go-interpreter/ast"
)

func TestLetStatements(t *testing.T) {
	const SENTENCES = 3
	input := `
	return x = 5;
	return y = 10;
	return 838383;
	`

	l := lexer.New(input)
	parser := New(l)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != SENTENCES {
		t.Fatalf("program.Statements does not contain %d statements. got=%d",
			SENTENCES, len(program.Statements))
	}

	// LET STATEMENT TESTS
	// tests := []struct {
	// 	expectedIdentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"foobar"},
	// }

	// for i, tt := range tests {
	// 	statement := program.Statements[i]
	// 	if !testLetStatement(t, statement, tt.expectedIdentifier) {
	// 		return
	// 	}
	// }
	
	// RETURN STATEMENT TESTS
	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not of type *ast.ReturnStatement, got=%T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got %q", 
				returnStatement.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, identifier string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	letStmt, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", statement)
	}

	if letStmt.Name.Value != identifier {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", identifier, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != identifier {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			identifier, letStmt.Name.TokenLiteral())
		return false
	}

	return true;
		
}

func checkParserErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	if len(errors) == 0 {
		return
	}

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.Fatalf("parser has %d errors", len(errors))
}