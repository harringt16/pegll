// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/goccmack/goutil/md"

	"github.com/bruceiv/pegll/test/lex/lex1/token"
)

type state int

const nullState state = -1

type Lexer struct {
	I      []rune
	Tokens []*token.Token
}

func NewFile(fname string) *Lexer {
	if strings.HasSuffix(fname, ".md") {
		src, err := md.GetSource(fname)
		if err != nil {
			panic(err)
		}
		return New([]rune(src))
	}
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return New([]rune(string(buf)))
}

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
			lext = tok.Rext
			lex.addToken(tok)
		}
	}
	lex.add(token.EOF, len(input), len(input))
	return lex
}

func (l *Lexer) scan(i int) *token.Token {
	// fmt.Printf("lexer.scan\n")
	s, tok := state(0), token.New(token.Error, i, i, nil)
	for s != nullState {
		// if tok.Rext >= len(l.I) {
		// fmt.Printf(" scan: state=%d tok=%s \"%s\" r=EOF\n",
		// s, tok, string(l.I[tok.Lext:tok.Rext]))
		// } else {
		// fmt.Printf(" scan: state=%d tok=%s \"%s\" r='%s'\n",
		// s, tok, string(l.I[tok.Lext:tok.Rext]), escape(l.I[tok.Rext]))
		// }
		if tok.Rext >= len(l.I) {
			s = nullState
		} else {
			tok.Type = accept[s]
			s = nextState[s](l.I[tok.Rext])
			if s != nullState || tok.Type == token.Error {
				tok.Rext++
			}
		}
	}
	tok.Literal = l.I[tok.Lext:tok.Rext]
	// fmt.Printf(" scan: state=%d tok=%s\n", s, tok)
	return tok
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

func (l *Lexer) GetLineColumnOfToken(i int) (line, col int) {
	return l.GetLineColumn(l.Tokens[i].Lext)
}

// GetString returns the input string from the left extent of Token[lext] to
// the right extent of Token[rext]
func (l *Lexer) GetString(lext, rext int) string {
	return string(l.I[l.Tokens[lext].Lext:l.Tokens[rext].Rext])
}

func (l *Lexer) add(t token.Type, lext, rext int) {
	l.addToken(token.New(t, lext, rext, l.I[lext:rext]))
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
	token.Error,
	token.Error,
	token.Type2,
}

var nextState = []func(r rune) state{
	// Set0
	func(r rune) state {
		switch {
		case r == 'a':
			return 1
		case r == 'b':
			return 2
		}
		return nullState
	},
	// Set1
	func(r rune) state {
		switch {
		case r == 'b':
			return 2
		}
		return nullState
	},
	// Set2
	func(r rune) state {
		switch {
		}
		return nullState
	},
}
