package parser

import (
  "baboon/lexer"
  "baboon/ast"
  "baboon/token"
)

type Parser struct {
  l *lexer.Lexer
  currToken token.Token
  nextToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{l: l}
  p.advance()
  p.advance()
  return p
}

func (p *Parser) advance() {
  p.currToken = p.nextToken
  p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  stmts := []ast.Statement{}

  for p.currToken.Type != token.EOF {
    if stmt := p.parseStatement; stmt != nil {
      stmts = append(stmts, p.parseStatement())
    }
    p.advance()
  }

  return &ast.Program{stmts}
}

func (p *Parser) parseStatement() ast.Statement {
  switch p.currToken.Type {
  case token.LET:
    return p.parseLetStatement()
  default:
    return nil
  }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
  stmt := &ast.LetStatement{Token: p.currToken}

  if (p.nextToken.Type == token.IDENT) {
    p.advance()
  } else {
    return nil
  }

  stmt.Name = ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

  if (p.nextToken.Type == token.ASSIGN) {
    p.advance()

  } else {
    return nil
  }

  for p.currToken.Type != token.SEMICOLON {
    p.advance()
  }

  return stmt
}
