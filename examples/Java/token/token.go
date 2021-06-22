
// Package token is generated by GoGLL. Do not edit
package token

import(
    "fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
    typ        Type
    lext, rext int
    input      []rune
}

/*
New returns a new token.
lext is the left extent and rext the right extent of the token in the input.
input is the input slice scanned by the lexer.
*/
func New(t Type, lext, rext int, input []rune) *Token {
    return &Token{
        typ:   t,
        lext:  lext,
        rext:  rext,
        input: input,
    }
}

// GetLineColumn returns the line and column of the left extent of t
func (t *Token) GetLineColumn() (line, col int) {
    line, col = 1, 1
    for j := 0; j < t.lext; j++ {
        switch t.input[j] {
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

// GetInput returns the input from which t was parsed.
func (t *Token) GetInput() []rune {
    return t.input
}

// Lext returns the left extent of t
func (t *Token) Lext() int {
    return t.lext
}

// Literal returs the literal runes of t scanned by the lexer
func (t *Token) Literal() []rune {
    return t.input[t.lext:t.rext]
}

// LiteralString returns string(t.Literal())
func (t *Token) LiteralString() string {
    return string(t.Literal())
}

// Rext returns the right extent of t in the input
func (t *Token) Rext() int {
    return t.rext
}

func (t *Token) String() string {
    return fmt.Sprintf("%s (%d,%d) %s",
        t.TypeID(), t.lext, t.rext, t.LiteralString())
}

// Suppress returns true iff t is suppressed by the lexer
func (t *Token) Suppress() bool {
	return Suppress[t.typ]
}

// Type returns the token Type of t
func (t *Token) Type() Type {
    return t.typ
}

// TypeID returns the token Type ID of t. 
// This may be different from the literal of token t.
func (t *Token) TypeID() string {
    return t.Type().ID()
}

// Type is the token type
type Type int

func (t Type) String() string {
    return TypeToString[t]
}

// ID returns the token type ID of token Type t
func (t Type) ID() string {
    return TypeToID[t]
}


const(
    Error  Type = iota  // Error 
    EOF  // $ 
    T_0  // block_comment 
    T_1  // escCharSp 
    T_2  // line_comment 
    T_3  // newline 
)

var TypeToString = []string{ 
    "Error",
    "EOF",
    "T_0",
    "T_1",
    "T_2",
    "T_3",
}

var StringToType = map[string] Type { 
    "Error" : Error, 
    "EOF" : EOF, 
    "T_0" : T_0, 
    "T_1" : T_1, 
    "T_2" : T_2, 
    "T_3" : T_3, 
}

var TypeToID = []string { 
    "Error", 
    "$", 
    "block_comment", 
    "escCharSp", 
    "line_comment", 
    "newline", 
}

var Suppress = []bool { 
    false, 
    false, 
    true, 
    false, 
    true, 
    false, 
}

