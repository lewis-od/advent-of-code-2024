package parsing

import (
	"slices"
)

type OpType int

const (
	OpMul OpType = iota
	OpDo
	OpDont
)

type Operation struct {
	Type        OpType
	Left, Right int
}

type Parser struct {
	tokens     []Token
	current    int
	operations []Operation
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, operations: []Operation{}}
}

// ops -> operation*
// operation -> mul | imperative
// mul -> "mul" "(" NUMBER "," NUMBER ")"
// imperative -> ( "do" | "don't" ) "(" ")"
func (p *Parser) Parse() []Operation {
	for !p.isAtEnd() {
		p.ops()
	}
	return p.operations
}

func (p *Parser) ops() {
	for p.chomp(TokenMul, TokenDo, TokenDont) {
		keyword := p.previous()
		_, wasPresent := p.consume(TokenLeftParen)
		if !wasPresent {
			continue
		}

		if keyword.Type != TokenMul {
			// imperative
			_, wasPresent := p.consume(TokenRightParen)
			if !wasPresent {
				continue
			}

			if keyword.Type == TokenDo {
				p.operations = append(p.operations, Operation{Type: OpDo})
			} else {
				p.operations = append(p.operations, Operation{Type: OpDont})
			}
			continue
		}

		left, wasPresent := p.consume(TokenNumber)
		if !wasPresent {
			continue
		}

		_, wasPresent = p.consume(TokenComma)
		if !wasPresent {
			continue
		}

		right, wasPresent := p.consume(TokenNumber)
		if !wasPresent {
			continue
		}

		_, wasPresent = p.consume(TokenRightParen)
		if !wasPresent {
			continue
		}

		operation := Operation{
			Type:  OpMul,
			Left:  left.Literal,
			Right: right.Literal,
		}
		p.operations = append(p.operations, operation)
	}
}

func (p *Parser) chomp(tokenTypes ...TokenType) bool {
	for !slices.Contains(tokenTypes, p.peek().Type) && !p.isAtEnd() {
		p.advance()
	}

	if p.isAtEnd() {
		// Didn't find the token we were after
		return false
	}
	p.advance()
	return true
}

func (p *Parser) consume(tokenType TokenType) (Token, bool) {
	if p.check(tokenType) {
		return p.advance(), true
	}
	return p.peek(), false
}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == TokenEof
}
