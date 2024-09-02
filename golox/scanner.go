package main

// Scanner tokenizes Lox source code. The implementation is largely based on the
// Go language scanner rather than the example code in the book.
type Scanner struct {
	src []byte

	offset     int // current read offset
	prevOffset int // first character of current lexeme being scanned
	lineOffset int // offset of first character of the current line
	lineNumber int // current line number, starting at 1

}

func NewScanner() *Scanner {
	var s Scanner
	s.reset()
	return &s
}

func (s *Scanner) Scan(src []byte) []Token {
	s.reset()
	s.src = src

	var tokens []Token
	for {
		t := s.scanToken()
		tokens = append(tokens, t)
		if t.Type == EOF {
			return tokens
		}
	}
}

func (s *Scanner) scanToken() (tok Token) {
	s.skipWhiteSpace()

	if s.eof() {
		tok.Type = EOF
		return tok
	}

	var t TokenType

	ch := rune(s.src[s.offset])
	if isLetter(ch) {
		lit := s.scanIdentifier()
		t = Lookup(lit)
		if t == Identifier {
			tok.Literal = lit
		}
	} else {
		switch ch {
		case '(':
			t = LeftParen
		case ')':
			t = RightParen
		case '{':
			t = LeftBrace
		case '}':
			t = RightBrace
		case ',':
			t = Comma
		case '.':
			t = Dot
		case '-':
			t = Minus
		case '+':
			t = Plus
		case ';':
			t = Semicolon
		case '*':
			t = Star
		// two-char tokens
		case '!':
			t = s.matchNext('=', BangEqual, Bang)
		case '=':
			t = s.matchNext('=', EqualEqual, Equal)
		case '<':
			t = s.matchNext('=', LessEqual, Less)
		case '>':
			t = s.matchNext('=', GreaterEqual, Greater)
		// slash
		case '/':
			if next, ok := s.peekNext(); ok && next == '/' {
				t = Slash
				break
			}
			// comments last until a newline
			for next, ok := s.peekNext(); ok && next != '\n'; {
				s.next()
			}
		default:
			t = Illegal
		}
	}

	tok.Type = t
	s.next()
	return tok
}

func (s *Scanner) reset() {
	s.offset = 0
	s.prevOffset = 0
	s.lineOffset = 0
	s.lineNumber = 1
}

// read the next ascii char
func (s *Scanner) next() {
	if !s.eof() {
		if s.src[s.offset] == '\n' {
			s.lineOffset = s.offset + 1
			s.lineNumber += 1
		}
		s.offset += 1
	}
	s.prevOffset = s.offset
}

func (s *Scanner) skipWhiteSpace() {
	for !s.eof() {
		switch s.src[s.offset] {
		case '\n':
		case ' ':
		case '\r':
		case '\t':
		default:
			return
		}

		s.next()
	}
}

func (s *Scanner) eof() bool {
	return s.offset >= len(s.src)
}

func (s *Scanner) matchNext(char rune, ifTrue TokenType, ifFalse TokenType) TokenType {
	if ch, ok := s.peekNext(); ok && ch == char {
		s.offset += 1
		return ifTrue
	}
	return ifFalse
}

func (s *Scanner) peekNext() (char rune, ok bool) {
	if s.offset+1 < len(s.src) {
		return rune(s.src[s.offset+1]), true
	}
	return ' ', false
}

// scanIdentifier reads the string of valid identifier characters at s.offset. It must
// only be called if it is known that the character at offset is a valid letter.
func (s *Scanner) scanIdentifier() string {
	startOffset := s.offset
	for n, b := range s.src[s.offset:] {
		// avoid rune conversion for subsequent characters
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
			continue
		}

		s.offset += n - 1
		s.next()
		return string(s.src[startOffset:s.offset])
	}
	s.offset = len(s.src)
	s.prevOffset = len(s.src)
	return string(s.src[startOffset:s.offset])
}

// return true if rune is an alphabetic character or underscore
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_'
}

func isDecimal(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// return lowercase ch iff ch is ASCII letter
func lower(ch rune) rune {
	return ('a' - 'A') | ch
}
