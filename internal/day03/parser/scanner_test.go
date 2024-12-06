package parser_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day03/parser"
	"github.com/stretchr/testify/require"
)

func TestScanMulOperationOnly(t *testing.T) {
	source := "mul(21,5)"

	scanner := parser.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parser.Token{
		{Type: parser.TokenMul, Lexeme: "mul"},
		{Type: parser.TokenLeftParen, Lexeme: "("},
		{Type: parser.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parser.TokenComma, Lexeme: ","},
		{Type: parser.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parser.TokenRightParen, Lexeme: ")"},
		{Type: parser.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}

func TestScanMulOperationWithSpace(t *testing.T) {
	source := "mul(21, 5)"

	scanner := parser.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parser.Token{
		{Type: parser.TokenMul, Lexeme: "mul"},
		{Type: parser.TokenLeftParen, Lexeme: "("},
		{Type: parser.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parser.TokenComma, Lexeme: ","},
		{Type: parser.TokenGarbage, Lexeme: " "},
		{Type: parser.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parser.TokenRightParen, Lexeme: ")"},
		{Type: parser.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}

func TestScanMulOperationsWithGarbage(t *testing.T) {
	source := "then(mul(11,8)mul(8,5))"

	scanner := parser.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parser.Token{
		{Type: parser.TokenGarbage, Lexeme: "then"},
		{Type: parser.TokenLeftParen, Lexeme: "("},
		{Type: parser.TokenMul, Lexeme: "mul"},
		{Type: parser.TokenLeftParen, Lexeme: "("},
		{Type: parser.TokenNumber, Lexeme: "11", Literal: 11},
		{Type: parser.TokenComma, Lexeme: ","},
		{Type: parser.TokenNumber, Lexeme: "8", Literal: 8},
		{Type: parser.TokenRightParen, Lexeme: ")"},
		{Type: parser.TokenMul, Lexeme: "mul"},
		{Type: parser.TokenLeftParen, Lexeme: "("},
		{Type: parser.TokenNumber, Lexeme: "8", Literal: 8},
		{Type: parser.TokenComma, Lexeme: ","},
		{Type: parser.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parser.TokenRightParen, Lexeme: ")"},
		{Type: parser.TokenRightParen, Lexeme: ")"},
		{Type: parser.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}
