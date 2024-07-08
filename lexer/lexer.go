package lexer

import (
  "baboon/token"
)

type Lexer struct {
  input string
  currentPosition int
  nextPosition int
  ch byte
}

func New(input string) *Lexer {
  ret := &Lexer{input: input}
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
    if isLetter(l.ch) {
      startPosition := l.currentPosition

      for isLetter(l.ch) {
        l.readChar()
      }
      literal := l.input[startPosition:l.currentPosition]
      return token.Token{token.LookupIdent(literal), literal}
    } else if isDigit(l.ch) {
      startPosition := l.currentPosition

      for isDigit(l.ch) {
        l.readChar()
      }
      literal := l.input[startPosition:l.currentPosition]
      return token.Token{token.INT, literal}
    } else {
      tok = token.Token{token.ILLEGAL, ""}
    } 
  }

  l.readChar()
  return tok
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}
func isLetter(ch byte) bool {
  return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
  if l.nextPosition >= len(l.input) {
    return 0
  } else {
    return l.input[l.nextPosition]
  }
}
