package lexer

import (
  "testing"
  "baboon/token"
)

func TestSimple(t *testing.T) {
  input := `=+(){},;-!/*<>`

  tests := []token.Token{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.MINUS, "-"},
    {token.BANG, "!"},
    {token.SLASH, "/"},
    {token.ASTERISK, "*"},
    {token.LT, "<"},
    {token.GT, ">"},
    {token.EOF, ""},
  }

  l := New(input)

  for i, expected := range tests {
    token := l.NextToken()

    if token.Type != expected.Type {
      t.Errorf("tests[%d]: Expected token type %q but was %q instead.", i, expected.Type, token.Type)
    }

    if token.Literal != expected.Literal {
      t.Errorf("tests[%d]: Expected literal %q but was %q instead.", i, expected.Literal, token.Literal)
    }

  }
}

func TestAssignment(t *testing.T) {
  input := `let x = 5 + 2;
            let camelCaseVariable = x + x;`

  tests := []token.Token{
    {token.LET, "let"},
    {token.IDENT, "x"},
    {token.ASSIGN, "="},
    {token.INT, "5"},
    {token.PLUS, "+"},
    {token.INT, "2"},
    {token.SEMICOLON, ";"},
    {token.LET, "let"},
    {token.IDENT, "camelCaseVariable"},
    {token.ASSIGN, "="},
    {token.IDENT, "x"},
    {token.PLUS, "+"},
    {token.IDENT, "x"},
    {token.SEMICOLON, ";"},
    {token.EOF, ""},
  }

  l := New(input)

  for i, expected := range tests {
    token := l.NextToken()

    if token.Type != expected.Type {
      t.Errorf("tests[%d]: Expected token type %q but was %q instead.", i, expected.Type, token.Type)
    }

    if token.Literal != expected.Literal {
      t.Errorf("tests[%d]: Expected literal %q but was %q instead.", i, expected.Literal, token.Literal)
    }
  }
}

func TestIsLetter(t *testing.T) {
  letters := "This_String_Is_OK"

  for i := 0; i < len(letters); i++ {
    if !isLetter(letters[i]) {
      t.Fatalf("Expected isLetter to return true but was false for %q.", letters[i])
    }
  }
}

func TestFromBook(t *testing.T) {
	input := `let five = 5;
            let ten = 10;

            let add = fn(x, y) {
              x + y;
            };

            let result = add(five, ten);
            !-/*5;
            5 < 10 > 5;

            if (5 < 10) {
              return true;
            } else {
              return false;
            }

            10 == 10;
            10 != 9;
            `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
