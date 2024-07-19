package ast

import (
  "testing"
  "baboon/token"
)

func TestProgramTokenLiteral(t *testing.T) {

  program := Program{}

  if program.TokenLiteral() != "" {
    t.Fatalf("TokenLiteral for empty Program wrong. Expected empty string, got %s instead.", program.TokenLiteral())
  }

  stmts := []Statement{
    &LetStatement{token.Token{token.LET, "let"}, Identifier{token.Token{token.IDENT, "x"}, "x"}},
  }
  program.Statements = stmts

  if program.TokenLiteral() != "let" {
    t.Fatalf("TokenLiteral for Program wrong. Expected \"let\" string, got %q instead.", program.TokenLiteral())
  }
}
