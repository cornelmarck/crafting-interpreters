package token

import (
	"errors"
	"strconv"
)

// Scanner tokenizes Lox source code. The implementation is based on the
// Go language scanner (go/scanner)
type Scanner struct {
	src []byte

	offset     int // current read offset
	prevOffset int // first character of current lexeme being scanned
	lineOffset int // offset of first character of the current line
	lineNumber int // current line number, starting at 1
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		src:        src,
		lineNumber: 1,
	}
}

func (s *Scanner) Scan() []Token {
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

	var t Type

	ch := rune(s.src[s.offset])
	if isLetter(ch) {
		lit := s.scanIdentifier()
		t = Lookup(lit)
		if t == Identifier {
			tok.Literal = lit
		}
	} else if isDecimal(ch) {
		lit, err := s.scanNumber()
		if err == nil {
			t = Number
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
		case '"':
			t = String
			lit, err := s.scanString()
			if err != nil {
				tok.Type = Illegal
				break
			}
			tok.Type = String
			tok.Literal = lit
		default:
			t = Illegal
		}
	}

	tok.Type = t
	s.next()
	return tok
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
		ch := s.src[s.offset]
		if ch == '\n' || ch == ' ' || ch == '\r' || ch == '\t' {
			s.next()
			continue
		}
		return
	}
}

func (s *Scanner) eof() bool {
	return s.offset >= len(s.src)
}

func (s *Scanner) matchNext(char rune, ifTrue Type, ifFalse Type) Type {
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

func (s *Scanner) scanString() (string, error) {
	// opening " is already consumed
	s.next()
	for {
		if s.offset == len(s.src) || s.src[s.offset] == '\n' {
			return "", errors.New("string literal not terminated")
		}

		if s.src[s.offset] == '"' {
			break
		}
		// we can increment without using next() because newlines are illegal
		s.offset += 1
	}

	lit := string(s.src[s.prevOffset:s.offset])
	return lit, nil
}

func (s *Scanner) scanNumber() (float64, error) {
	if s.offset == len(s.src)-1 {
		literal := string(s.src[s.prevOffset:s.offset])
		return strconv.ParseFloat(literal, 64)
	}

	for s.offset+1 < len(s.src) {
		ch := rune(s.src[s.offset+1])
		if !isDecimal(ch) && ch != '.' {
			break
		}
		s.offset += 1
	}
	literal := string(s.src[s.prevOffset : s.offset+1])
	s.prevOffset = s.offset
	return strconv.ParseFloat(literal, 64)
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
