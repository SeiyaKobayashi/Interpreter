package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"testing"
)

// Test case
func TestLetStatements(t *testing.T) {
	input := `
    let x = 3;
    let y = 12;
    let foo = 1225;
    `

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	if prog == nil {
		t.Fatalf("ParseProgram() returned nil.")
	}
	if len(prog.Statements) != 3 {
		t.Fatalf("prog.statements doesn't contain 3 statements. got=%d",
			len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q",
			s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not ''%s'. got=%s",
			name,
			letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name,
			letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
