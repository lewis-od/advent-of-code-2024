package parsing_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day03/parsing"
	"github.com/stretchr/testify/require"
)

func TestParseMulOperation(t *testing.T) {
	mulTokens := []parsing.Token{
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(mulTokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{
		{Type: parsing.OpMul, Left: 21, Right: 5},
	}
	require.Equal(t, expectedOperations, operations)
}

func TestParseConsecutiveMulOperations(t *testing.T) {
	mulTokens := []parsing.Token{
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "4", Literal: 4},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "91", Literal: 91},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(mulTokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{
		{Type: parsing.OpMul, Left: 21, Right: 5},
		{Type: parsing.OpMul, Left: 4, Right: 91},
	}
	require.Equal(t, expectedOperations, operations)
}

func TestParseMulOperationWithGarbageSurrounding(t *testing.T) {
	tokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenMul, Lexeme: "mul"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenNumber, Lexeme: "21", Literal: 21},
		{Type: parsing.TokenComma, Lexeme: ","},
		{Type: parsing.TokenNumber, Lexeme: "5", Literal: 5},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenGarbage, Lexeme: "ignore me"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(tokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{
		{Type: parsing.OpMul, Left: 21, Right: 5},
	}
	require.Equal(t, expectedOperations, operations)
}

func TestParseDo(t *testing.T) {
	tokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenDo, Lexeme: "do"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenGarbage, Lexeme: "ignore me"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(tokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{
		{Type: parsing.OpDo},
	}
	require.Equal(t, expectedOperations, operations)
}

func TestParseInvalidDo(t *testing.T) {
	tokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenDo, Lexeme: "do"},
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenGarbage, Lexeme: "ignore me"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(tokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{}
	require.Equal(t, expectedOperations, operations)
}

func TestParseDont(t *testing.T) {
	tokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenDont, Lexeme: "dont"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenGarbage, Lexeme: "ignore me"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(tokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{
		{Type: parsing.OpDont},
	}
	require.Equal(t, expectedOperations, operations)
}

func TestParseInvalidDont(t *testing.T) {
	tokens := []parsing.Token{
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenDont, Lexeme: "dont"},
		{Type: parsing.TokenLeftParen, Lexeme: "("},
		{Type: parsing.TokenGarbage, Lexeme: "thing"},
		{Type: parsing.TokenRightParen, Lexeme: ")"},
		{Type: parsing.TokenGarbage, Lexeme: "ignore me"},
		{Type: parsing.TokenEof},
	}

	p := parsing.NewParser(tokens)
	operations := p.Parse()

	expectedOperations := []parsing.Operation{}
	require.Equal(t, expectedOperations, operations)
}
