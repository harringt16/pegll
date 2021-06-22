
// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface{
	isSymbol()
	IsNonTerminal() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol() {}

// NT is the type of non-terminals symbols
type NT int
const( 
	NT_LineOrBlock NT = iota
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // block_comment 
	T_1  // line_comment 
)

type Symbols []Symbol

func (ss Symbols) Strings() []string {
	strs := make([]string, len(ss))
	for i, s := range ss {
		strs[i] = s.String()
	}
	return strs
}

func (NT) IsNonTerminal() bool {
	return true
}

func (T) IsNonTerminal() bool {
	return false
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

var ntToString = []string { 
	"LineOrBlock", /* NT_LineOrBlock */ 
}

var tToString = []string { 
	"block_comment", /* T_0 */
	"line_comment", /* T_1 */ 
}

var stringNT = map[string]NT{ 
	"LineOrBlock":NT_LineOrBlock,
}
