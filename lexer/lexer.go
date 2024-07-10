package lexer

import (
  "baboon/token"
)

type Lexer struct {
  input []rune
  currentPosition int
  nextPosition int
  ch rune
}

func New(input string) *Lexer {
  ret := &Lexer{input: []rune(input)}
  ret.readChar()
  return ret
}

func (l *Lexer) readChar() {
  if l.nextPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.nextPosition]
  }
  l.currentPosition = l.nextPosition
  l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {

  var tok token.Token

  l.skipWhitespace()

  switch l.ch {
  case '+':
    tok = token.Token{token.PLUS, string(l.ch)}
  case ';':
    tok = token.Token{token.SEMICOLON, string(l.ch)}
  case '(':
    tok = token.Token{token.LPAREN, string(l.ch)}
  case ')':
    tok = token.Token{token.RPAREN, string(l.ch)}
  case '{':
    tok = token.Token{token.LBRACE, string(l.ch)}
  case '}':
    tok = token.Token{token.RBRACE, string(l.ch)}
  case ',':
    tok = token.Token{token.COMMA, string(l.ch)}
  case '/':
    tok = token.Token{token.SLASH, string(l.ch)}
  case '-':
    tok = token.Token{token.MINUS, string(l.ch)}
  case '<':
    tok = token.Token{token.LT, string(l.ch)}
  case '>':
    tok = token.Token{token.GT, string(l.ch)}
  case '*':
    tok = token.Token{token.ASTERISK, string(l.ch)}
  case '=':
    if l.peekChar() == '=' {
      tok = token.Token{token.EQ, "=="}
      l.readChar()
    } else {
      tok = token.Token{token.ASSIGN, string(l.ch)}
    }
  case '!':
    if l.peekChar() == '=' {
      tok = token.Token{token.NOT_EQ, "!="}
      l.readChar()
    } else {
      tok = token.Token{token.BANG, string(l.ch)}
    }
  case 0:
    tok = token.Token{token.EOF, ""}
  default:
    if isDigit(l.ch) {
      return token.Token{token.INT, readLiteral(l, isDigit)}
    } else {
      literal := readLiteral(l, func(ch rune) bool {
        return isLetter(ch) || (!token.IsDelimiter(ch) && !token.IsOperator(ch) && !isWhitespace(ch))
      })
      return token.Token{token.LookupIdent(literal), literal}
    } 
  }

  l.readChar()
  return tok
}

func readLiteral(l *Lexer, predicate func(rune) bool) string {
  startPosition := l.currentPosition

  for predicate(l.ch) {
    l.readChar()
  }
  return string(l.input[startPosition:l.currentPosition])
}

func (l *Lexer) skipWhitespace() {
  for isWhitespace(l.ch) {
    l.readChar()
  }
}

func isWhitespace(ch rune) bool {
  return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
  return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() rune {
  if l.nextPosition >= len(l.input) {
    return 0
  } else {
    return l.input[l.nextPosition]
  }
}
