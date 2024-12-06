package parsing

type OpType int

const (
	OpMul OpType = iota
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
	return &Parser{tokens: tokens}
}

// ops -> operation*
// operation -> mul
// mul -> "mul" "(" NUMBER "," NUMBER ")"
func (p *Parser) Parse() []Operation {
	for !p.isAtEnd() {
		p.ops()
	}
	return p.operations
}

func (p *Parser) ops() {
	for p.chomp(TokenMul) {
		_, wasPresent := p.consume(TokenLeftParen)
		if !wasPresent {
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

func (p *Parser) chomp(tokenType TokenType) bool {
	for p.peek().Type != tokenType && !p.isAtEnd() {
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
