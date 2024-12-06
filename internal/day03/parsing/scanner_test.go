package parsing_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day03/parsing"
	"github.com/stretchr/testify/require"
)

func TestScanMulOperationOnly(t *testing.T) {
	source := "mul(21,5)"

	scanner := parsing.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parsing.Token{
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}

func TestScanMulOperationWithSpace(t *testing.T) {
	source := "mul(21, 5)"

	scanner := parsing.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parsing.Token{
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenGarbage, Lexeme: " "},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}

func TestScanMulOperationsWithGarbage(t *testing.T) {
	source := "then(mul(11,8)mul(8,5))"

	scanner := parsing.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "then"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "11", Literal: 11},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "8", Literal: 8},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "8", Literal: 8},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}

func TestScanWithWordAttachedToMul(t *testing.T) {
	source := "xmul(1,2)"

	scanner := parsing.NewScanner(source)
	tokens := scanner.ScanTokens()

	expectedTokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "x"},
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "1", Literal: 1},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "2", Literal: 2},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}
	require.Equal(t, expectedTokens, tokens)
}
