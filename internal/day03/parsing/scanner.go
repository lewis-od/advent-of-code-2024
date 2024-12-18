package parsing

import (
	"strconv"
	"unicode"

	"github.com/lewis-od/aoc24/internal/common"
)

type TokenType int

const (
	TokenMul TokenType = iota
	TokenLeftParen
	TokenRightParen
	TokenComma
	TokenNumber
	TokenGarbage
	TokenDont // TODO
	TokenDo   // TODO
	TokenEof
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal int
}

type Scanner interface {
	ScanTokens() []Token
}

type scanner struct {
	source         string
	start, current int
	tokens         []Token
}

func NewScanner(source string) Scanner {
	return &scanner{source: source}
}

func (s *scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{Type: TokenEof})
	return s.tokens
}

func (s *scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(TokenLeftParen)
	case ')':
		s.addToken(TokenRightParen)
	case ',':
		s.addToken(TokenComma)
	default:
		if unicode.IsDigit(c) {
			s.number()
		} else if unicode.IsLetter(c) {
			s.keywordOrGarbage()
		} else {
			s.addToken(TokenGarbage)
		}
	}
}

func (s *scanner) number() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	lexeme := s.source[s.start:s.current]
	literal := common.Must(strconv.Atoi(lexeme))
	s.addNumberToken(literal)
}

func (s *scanner) keywordOrGarbage() {
	next := s.peek()
	lexeme := s.source[s.start:s.current]
	for unicode.IsLetter(next) && next != 'm' && next != 'd' && lexeme != "do" {
		s.advance()
		next = s.peek()
		lexeme = s.source[s.start:s.current]
	}

	if lexeme == "mul" {
		s.addToken(TokenMul)
		return
	} else if lexeme == "do" {
		if !s.match("n't") {
			s.addToken(TokenDo)
			return
		}

		s.dontOrGarbage()
		return
	}
	s.addToken(TokenGarbage)
}

func (s *scanner) match(lexeme string) bool {
	if s.current+len(lexeme) >= len(s.source) {
		return false
	}

	for i, char := range lexeme {
		if rune(s.source[s.current+i]) != char {
			return false
		}
	}
	return true
}

func (s *scanner) dontOrGarbage() {
	s.advance() // Consume 'n'
	first := s.advance()
	if first != '\'' {
		s.addToken(TokenGarbage)
		return
	}

	second := s.advance()
	if second != 't' {
		s.addToken(TokenGarbage)
		return
	}

	s.addToken(TokenDont)
}

func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return '\n'
	}
	return rune(s.source[s.current])
}

func (s *scanner) addToken(tokenType TokenType) {
	token := Token{Type: tokenType, Lexeme: s.source[s.start:s.current]}
	s.tokens = append(s.tokens, token)
}

func (s *scanner) addNumberToken(literal int) {
	token := Token{
		Type:    TokenNumber,
		Lexeme:  s.source[s.start:s.current],
		Literal: literal,
	}
	s.tokens = append(s.tokens, token)
}

func (s *scanner) advance() rune {
	char := rune(s.source[s.current])
	s.current++
	return char
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
