package parser

import (
  "testing"
  "baboon/lexer"
  "baboon/ast"
  "baboon/token"
)

func TestParseProgram(t *testing.T) {
  input := `
  let x = 123;
  let y = 321;
  let foobar = 444;
  `

  l := lexer.New(input)

  p := New(l)
  program := p.ParseProgram()

  if len(program.Statements) != 3 {
    t.Fatalf("Expected 3 statements to be parsed, but was %d.", len(program.Statements))
  }

  tests := []string{
    "x",
    "y",
    "foobar",
  }

  for i, tt := range tests {
    s := program.Statements[i]
    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
      t.Fatalf("Statement not *ast.LetStatement but %T", s)
    }
    if  letStmt.Name.Value != tt {
      t.Fatalf("Expected %s but was %s instead", tt, letStmt.Name.Value)
    }
    if letStmt.Token.Type != token.LET {
      t.Fatalf("letStmt does not have token.LET but %s instead.", letStmt.TokenLiteral())
    }
  }

}
