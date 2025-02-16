
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
	NT_CLOSE NT = iota
	NT_DIVIDE 
	NT_ELEMENT 
	NT_EXPR 
	NT_MINUS 
	NT_Number 
	NT_OPEN 
	NT_PLUS 
	NT_PLUSorMINUS 
	NT_PRODUCT 
	NT_RepPLUSorMINUS0x 
	NT_RepTIMESorDIV0x 
	NT_SUM 
	NT_TIMES 
	NT_TIMESorDIVIDE 
	NT_WS 
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // ( 
	T_1  // ) 
	T_2  // * 
	T_3  // + 
	T_4  // - 
	T_5  // / 
	T_6  // repNumber1x 
	T_7  // sp 
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
	"CLOSE", /* NT_CLOSE */
	"DIVIDE", /* NT_DIVIDE */
	"ELEMENT", /* NT_ELEMENT */
	"EXPR", /* NT_EXPR */
	"MINUS", /* NT_MINUS */
	"Number", /* NT_Number */
	"OPEN", /* NT_OPEN */
	"PLUS", /* NT_PLUS */
	"PLUSorMINUS", /* NT_PLUSorMINUS */
	"PRODUCT", /* NT_PRODUCT */
	"RepPLUSorMINUS0x", /* NT_RepPLUSorMINUS0x */
	"RepTIMESorDIV0x", /* NT_RepTIMESorDIV0x */
	"SUM", /* NT_SUM */
	"TIMES", /* NT_TIMES */
	"TIMESorDIVIDE", /* NT_TIMESorDIVIDE */
	"WS", /* NT_WS */ 
}

var tToString = []string { 
	"(", /* T_0 */
	")", /* T_1 */
	"*", /* T_2 */
	"+", /* T_3 */
	"-", /* T_4 */
	"/", /* T_5 */
	"repNumber1x", /* T_6 */
	"sp", /* T_7 */ 
}

var stringNT = map[string]NT{ 
	"CLOSE":NT_CLOSE,
	"DIVIDE":NT_DIVIDE,
	"ELEMENT":NT_ELEMENT,
	"EXPR":NT_EXPR,
	"MINUS":NT_MINUS,
	"Number":NT_Number,
	"OPEN":NT_OPEN,
	"PLUS":NT_PLUS,
	"PLUSorMINUS":NT_PLUSorMINUS,
	"PRODUCT":NT_PRODUCT,
	"RepPLUSorMINUS0x":NT_RepPLUSorMINUS0x,
	"RepTIMESorDIV0x":NT_RepTIMESorDIV0x,
	"SUM":NT_SUM,
	"TIMES":NT_TIMES,
	"TIMESorDIVIDE":NT_TIMESorDIVIDE,
	"WS":NT_WS,
}
