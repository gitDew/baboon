package parser

import (
  "testing"
  "baboon/lexer"
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
}
