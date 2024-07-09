package token

import "slices"

type Token struct {
  Type    TokenType
  Literal string
}

type TokenType string

const (
    // Identifiers + Literals
    IDENT       = "IDENT"
    INT         = "INT"

    // Operators
    ASSIGN      = "="
    PLUS        = "+"
    MINUS       = "-"     
    BANG        = "!"
    SLASH       = "/"
    ASTERISK    = "*"
    LT          = "<"
    GT          = ">"
    EQ          = "=="
    NOT_EQ      = "!="


    // Delimiters
    SEMICOLON   = ";"
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"
    COMMA       = ","

    // Keywords
    LET         = "LET"
    FUNCTION    = "FUNCTION"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"

    EOF         = "EOF"
    ILLEGAL     = "ILLEGAL"
)

var keywords = map[string]TokenType{
    "let":      LET,
    "fn":       FUNCTION,
    "true":     TRUE,
    "false":    FALSE,
    "if":       IF,
    "else":     ELSE,
    "return":   RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

var delimiterChars = []rune{
    ';',
    '(',
    ')',
    '{',
    '}',
    ',',
}

var operatorChars = []rune{
    '=',
    '+',
    '-',
    '!',
    '/',
    '*',
    '<',
    '>',
}

func IsOperator(ch rune) bool {
    return slices.Contains(operatorChars, ch)
}

func IsDelimiter(ch rune) bool {
    return slices.Contains(delimiterChars, ch)
}
