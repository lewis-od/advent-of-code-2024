package parser

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
	TokenEof
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal int
}

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, line: 1}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{Type: TokenEof})
	return s.tokens
}

func (s *Scanner) scanToken() {
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
			s.keyword()
		} else {
			s.addToken(TokenGarbage)
		}
	}
}

func (s *Scanner) number() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	lexeme := s.source[s.start:s.current]
	literal := common.Must(strconv.Atoi(lexeme))
	s.addNumberToken(literal)
}

func (s *Scanner) keyword() {
	for unicode.IsLetter(s.peek()) {
		s.advance()
	}

	lexeme := s.source[s.start:s.current]
	if lexeme == "mul" {
		s.addToken(TokenMul)
	} else {
		s.addToken(TokenGarbage)
	}
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\n'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) addToken(tokenType TokenType) {
	token := Token{Type: tokenType, Lexeme: s.source[s.start:s.current]}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) addNumberToken(literal int) {
	token := Token{
		Type:    TokenNumber,
		Lexeme:  s.source[s.start:s.current],
		Literal: literal,
	}
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) advance() rune {
	char := rune(s.source[s.current])
	s.current++
	return char
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
