package lexer

import (
	"github.com/vickychhetri/mon-ter/token"
)

type lexer struct {
	input       string
	position    int
	readPostion int
	ch          byte
}

func New(input string) *lexer {
	l := &lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *lexer) readChar() {

	if l.readPostion >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPostion]
	}

	l.position = l.readPostion
	l.readPostion += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()

	return tok

}
