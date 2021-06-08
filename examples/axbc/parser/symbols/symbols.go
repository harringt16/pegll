
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
	NT_AxBC NT = iota
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // ab 
	T_1  // as 
	T_2  // c 
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
	"AxBC", /* NT_AxBC */ 
}

var tToString = []string { 
	"ab", /* T_0 */
	"as", /* T_1 */
	"c", /* T_2 */ 
}

var stringNT = map[string]NT{ 
	"AxBC":NT_AxBC,
}
