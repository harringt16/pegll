
// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"calc/token"
)

type state int

const nullState state = -1


// Lexer contains both the input slice of runes and the slice of tokens
// parsed from the input
type Lexer struct {
	// I is the input slice of runes
	I      []rune

	// Tokens is the slice of tokens constructed by the lexer from I
	Tokens []*token.Token
}

/*
NewFile constructs a Lexer created from the input file, fname. 

If the input file is a markdown file NewFile process treats all text outside
code blocks as whitespace. All text inside code blocks are treated as input text.

If the input file is a normal text file NewFile treats all text in the inputfile
as input text.
*/
func NewFile(fname string) *Lexer {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	input := []rune(string(buf))
	if strings.HasSuffix(fname, ".md") {
		loadMd(input)
	}
	return New(input)
}

func loadMd(input []rune) {
	i := 0
	text := true
	for i < len(input) {
		if i <= len(input)-3 && input[i] == '`' && input[i+1] == '`' && input[i+2] == '`' {
			text = !text
			for j := 0; j < 3; j++ {
				input[i+j] = ' '
			}
			i += 3
		}
		if i < len(input) {
			if text {
				if input[i] == '\n' {
					input[i] = '\n'
				} else {
					input[i] = ' '
				}
			}
			i += 1
		}
	}
}

/*
New constructs a Lexer from a slice of runes. 

All contents of the input slice are treated as input text.
*/
func New(input []rune) *Lexer {
	lex := &Lexer{
		I:      input,
		Tokens: make([]*token.Token, 0, 2048),
	}
	lext := 0
	for lext < len(lex.I) {
		for lext < len(lex.I) && unicode.IsSpace(lex.I[lext]) {
			lext++
		}
		if lext < len(lex.I) {
			tok := lex.scan(lext)
			lext = tok.Rext()
			if !tok.Suppress() {
				lex.addToken(tok)
			}
		}
	}
	lex.add(token.EOF, len(input), len(input))
	return lex
}

func (l *Lexer) scan(i int) *token.Token {
	// fmt.Printf("lexer.scan\n")
	s, typ, rext := state(0), token.Error, i
	for s != nullState {
		// fmt.Printf("S%d '%c' @ %d\n", s, l.I[rext], rext)
		if rext >= len(l.I) {
			typ = accept[s]
			s = nullState
		} else {
			typ = accept[s]
			s = nextState[s](l.I[rext])
			if s != nullState || typ == token.Error {
				rext++
			}
		}
	}
	return token.New(typ, i, rext, l.I)
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetLineColumnOfToken returns the line and column of token[i] in the imput
func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext())
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext():l.Tokens[rext].Rext()])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.New(t, lext, rext, l.I))
}

func (l *Lexer) addToken(tok *token.Token) {
	l.Tokens = append(l.Tokens, tok)
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = []token.Type{ 
	token.T_7, 
	token.T_7, 
	token.T_0, 
	token.T_1, 
	token.T_2, 
	token.T_3, 
	token.T_4, 
	token.T_5, 
	token.T_6, 
}

var nextState = []func(r rune) state{ 
	// Set0
	func(r rune) state {
		switch { 
		case r == '\t':
			return 1 
		case r == ' ':
			return 1 
		case r == '(':
			return 2 
		case r == ')':
			return 3 
		case r == '*':
			return 4 
		case r == '+':
			return 5 
		case r == '-':
			return 6 
		case r == '/':
			return 7 
		case unicode.IsNumber(r):
			return 8 
		}
		return nullState
	}, 
	// Set1
	func(r rune) state {
		switch { 
		case r == '\t':
			return 1 
		case r == ' ':
			return 1 
		}
		return nullState
	}, 
	// Set2
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set3
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set4
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set5
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set6
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set7
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set8
	func(r rune) state {
		switch { 
		case unicode.IsNumber(r):
			return 8 
		}
		return nullState
	}, 
}
