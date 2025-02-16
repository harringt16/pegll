// Package parser is generated by gogll. Do not edit.
package parser

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"axbc/lexer"
	"axbc/parser/bsr"
	"axbc/parser/slot"
	"axbc/parser/symbols"
	"axbc/token"
)

type parser struct {
	cI int

	R *descriptors
	U *descriptors

	popped   map[poppedNode]bool
	crf      map[clusterNode][]*crfNode
	crfNodes map[crfNode]*crfNode

	lex         *lexer.Lexer
	parseErrors []*Error

	bsrSet *bsr.Set
}

func newParser(l *lexer.Lexer) *parser {
	return &parser{
		cI:     0,
		lex:    l,
		R:      &descriptors{},
		U:      &descriptors{},
		popped: make(map[poppedNode]bool),
		crf: map[clusterNode][]*crfNode{
			{symbols.NT_AxBC, 0}: {},
		},
		crfNodes:    map[crfNode]*crfNode{},
		bsrSet:      bsr.New(symbols.NT_AxBC, l),
		parseErrors: nil,
	}
}

// Parse returns the BSR set containing the parse forest.
// If the parse was successfull []*Error is nil
func Parse(l *lexer.Lexer) (*bsr.Set, []*Error) {
	return newParser(l).parse()
}

func (p *parser) parse() (*bsr.Set, []*Error) {
	var L slot.Label
	m, cU := len(p.lex.Tokens)-1, 0
	p.ntAdd(symbols.NT_AxBC, 0)
	// p.DumpDescriptors()
	for !p.R.empty() {
		L, cU, p.cI = p.R.remove()

		// fmt.Println()
		// fmt.Printf("L:%s, cI:%d, I[p.cI]:%s, cU:%d\n", L, p.cI, p.lex.Tokens[p.cI], cU)
		// p.DumpDescriptors()

		switch L {
		case slot.AorB0R0: // AorB : ∙Repa0x

			p.call(slot.AorB0R1, cU, p.cI)
		case slot.AorB0R1: // AorB : Repa0x ∙

			p.rtn(symbols.NT_AorB, cU, p.cI)
		case slot.AorB1R0: // AorB : ∙a b

			p.bsrSet.Add(slot.AorB1R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.AorB1R1) {
				p.parseError(slot.AorB1R1, p.cI, first[slot.AorB1R1])
				break
			}

			p.bsrSet.Add(slot.AorB1R2, cU, p.cI, p.cI+1)
			p.cI++
			p.rtn(symbols.NT_AorB, cU, p.cI)
		case slot.AxBC0R0: // AxBC : ∙AorB c

			p.call(slot.AxBC0R1, cU, p.cI)
		case slot.AxBC0R1: // AxBC : AorB ∙c

			if !p.testSelect(slot.AxBC0R1) {
				p.parseError(slot.AxBC0R1, p.cI, first[slot.AxBC0R1])
				break
			}

			p.bsrSet.Add(slot.AxBC0R2, cU, p.cI, p.cI+1)
			p.cI++
			p.rtn(symbols.NT_AxBC, cU, p.cI)
		case slot.Repa0x0R0: // Repa0x : ∙a Repa0x

			p.bsrSet.Add(slot.Repa0x0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.Repa0x0R1) {
				p.parseError(slot.Repa0x0R1, p.cI, first[slot.Repa0x0R1])
				break
			}

			p.call(slot.Repa0x0R2, cU, p.cI)
		case slot.Repa0x0R2: // Repa0x : a Repa0x ∙

			p.rtn(symbols.NT_Repa0x, cU, p.cI)
		case slot.Repa0x1R0: // Repa0x : ∙
			p.bsrSet.AddEmpty(slot.Repa0x1R0, p.cI)

			p.rtn(symbols.NT_Repa0x, cU, p.cI)

		default:
			panic("This must not happen")
		}
	}
	if !p.bsrSet.Contain(symbols.NT_AxBC, 0, m) {
		p.sortParseErrors()
		return nil, p.parseErrors
	}
	return p.bsrSet, nil
}

func (p *parser) ntAdd(nt symbols.NT, j int) {
	// fmt.Printf("p.ntAdd(%s, %d)\n", nt, j)
	failed := true
	expected := map[token.Type]string{}
	for _, l := range slot.GetAlternates(nt) {
		if p.testSelect(l) {
			p.dscAdd(l, j, j)
			failed = false
		} else {
			for k, v := range first[l] {
				expected[k] = v
			}
		}
	}
	if failed {
		for _, l := range slot.GetAlternates(nt) {
			p.parseError(l, j, expected)
		}
	}
}

/*** Call Return Forest ***/

type poppedNode struct {
	X    symbols.NT
	k, j int
}

type clusterNode struct {
	X symbols.NT
	k int
}

type crfNode struct {
	L slot.Label
	i int
}

/*
suppose that L is Y ::=αX ·β
if there is no CRF node labelled (L,i)
	create one let u be the CRF node labelled (L,i)
if there is no CRF node labelled (X, j) {
	create a CRF node v labelled (X, j)
	create an edge from v to u
	ntAdd(X, j)
} else {
	let v be the CRF node labelled (X, j)
	if there is not an edge from v to u {
		create an edge from v to u
		for all ((X, j,h)∈P) {
			dscAdd(L, i, h);
			bsrAdd(L, i, j, h)
		}
	}
}
*/
func (p *parser) call(L slot.Label, i, j int) {
	// fmt.Printf("p.call(%s,%d,%d)\n", L,i,j)
	u, exist := p.crfNodes[crfNode{L, i}]
	// fmt.Printf("  u exist=%t\n", exist)
	if !exist {
		u = &crfNode{L, i}
		p.crfNodes[*u] = u
	}
	X := L.Symbols()[L.Pos()-1].(symbols.NT)
	ndV := clusterNode{X, j}
	v, exist := p.crf[ndV]
	if !exist {
		// fmt.Println("  v !exist")
		p.crf[ndV] = []*crfNode{u}
		p.ntAdd(X, j)
	} else {
		// fmt.Println("  v exist")
		if !existEdge(v, u) {
			// fmt.Printf("  !existEdge(%v)\n", u)
			p.crf[ndV] = append(v, u)
			// fmt.Printf("|popped|=%d\n", len(popped))
			for pnd := range p.popped {
				if pnd.X == X && pnd.k == j {
					p.dscAdd(L, i, pnd.j)
					p.bsrSet.Add(L, i, j, pnd.j)
				}
			}
		}
	}
}

func existEdge(nds []*crfNode, nd *crfNode) bool {
	for _, nd1 := range nds {
		if nd1 == nd {
			return true
		}
	}
	return false
}

func (p *parser) rtn(X symbols.NT, k, j int) {
	// fmt.Printf("p.rtn(%s,%d,%d)\n", X,k,j)
	pn := poppedNode{X, k, j}
	if _, exist := p.popped[pn]; !exist {
		p.popped[pn] = true
		for _, nd := range p.crf[clusterNode{X, k}] {
			p.dscAdd(nd.L, nd.i, j)
			p.bsrSet.Add(nd.L, nd.i, k, j)
		}
	}
}

// func CRFString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("CRF: {")
// 	for cn, nds := range crf{
// 		for _, nd := range nds {
// 			fmt.Fprintf(buf, "%s->%s, ", cn, nd)
// 		}
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

func (cn clusterNode) String() string {
	return fmt.Sprintf("(%s,%d)", cn.X, cn.k)
}

func (n crfNode) String() string {
	return fmt.Sprintf("(%s,%d)", n.L.String(), n.i)
}

// func PoppedString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("Popped: {")
// 	for p, _ := range popped {
// 		fmt.Fprintf(buf, "(%s,%d,%d) ", p.X, p.k, p.j)
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

/*** descriptors ***/

type descriptors struct {
	set []*descriptor
}

func (ds *descriptors) contain(d *descriptor) bool {
	for _, d1 := range ds.set {
		if d1 == d {
			return true
		}
	}
	return false
}

func (ds *descriptors) empty() bool {
	return len(ds.set) == 0
}

func (ds *descriptors) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	for i, d := range ds.set {
		if i > 0 {
			buf.WriteString("; ")
		}
		fmt.Fprintf(buf, "%s", d)
	}
	buf.WriteString("}")
	return buf.String()
}

type descriptor struct {
	L slot.Label
	k int
	i int
}

func (d *descriptor) String() string {
	return fmt.Sprintf("%s,%d,%d", d.L, d.k, d.i)
}

func (p *parser) dscAdd(L slot.Label, k, i int) {
	// fmt.Printf("p.dscAdd(%s,%d,%d)\n", L, k, i)
	d := &descriptor{L, k, i}
	if !p.U.contain(d) {
		p.R.set = append(p.R.set, d)
		p.U.set = append(p.U.set, d)
	}
}

func (ds *descriptors) remove() (L slot.Label, k, i int) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	// fmt.Printf("remove: %s,%d,%d\n", d.L, d.k, d.i)
	return d.L, d.k, d.i
}

func (p *parser) DumpDescriptors() {
	p.DumpR()
	p.DumpU()
}

func (p *parser) DumpR() {
	fmt.Println("R:")
	for _, d := range p.R.set {
		fmt.Printf(" %s\n", d)
	}
}

func (p *parser) DumpU() {
	fmt.Println("U:")
	for _, d := range p.U.set {
		fmt.Printf(" %s\n", d)
	}
}

/*** TestSelect ***/

func (p *parser) follow(nt symbols.NT) bool {
	_, exist := followSets[nt][p.lex.Tokens[p.cI].Type()]
	return exist
}

func (p *parser) testSelect(l slot.Label) bool {
	return l.IsNullable() || l.FirstContains(p.lex.Tokens[p.cI].Type())
	// _, exist := first[l][p.lex.Tokens[p.cI].Type()]
	// return exist
}

var first = []map[token.Type]string{
	// AorB : ∙Repa0x
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// AorB : Repa0x ∙
	{
		token.T_2: "c",
	},
	// AorB : ∙a b
	{
		token.T_0: "a",
	},
	// AorB : a ∙b
	{
		token.T_1: "b",
	},
	// AorB : a b ∙
	{
		token.T_2: "c",
	},
	// AxBC : ∙AorB c
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// AxBC : AorB ∙c
	{
		token.T_2: "c",
	},
	// AxBC : AorB c ∙
	{
		token.EOF: "$",
	},
	// Repa0x : ∙a Repa0x
	{
		token.T_0: "a",
	},
	// Repa0x : a ∙Repa0x
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// Repa0x : a Repa0x ∙
	{
		token.T_2: "c",
	},
	// Repa0x : ∙
	{
		token.T_2: "c",
	},
}

var followSets = []map[token.Type]string{
	// AorB
	{
		token.T_2: "c",
	},
	// AxBC
	{
		token.EOF: "$",
	},
	// Repa0x
	{
		token.T_2: "c",
	},
}

/*** Errors ***/

/*
Error is returned by Parse at every point at which the parser fails to parse
a grammar production. For non-LL-1 grammars there will be an error for each
alternate attempted by the parser.

The errors are sorted in descending order of input position (index of token in
the stream of tokens).

Normally the error of interest is the one that has parsed the largest number of
tokens.
*/
type Error struct {
	// Index of token that caused the error.
	cI int

	// Grammar slot at which the error occured.
	Slot slot.Label

	// The token at which the error occurred.
	Token *token.Token

	// The line and column in the input text at which the error occurred
	Line, Column int

	// The tokens expected at the point where the error occurred
	Expected map[token.Type]string
}

func (pe *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Parse Error: %s I[%d]=%s at line %d col %d\n",
		pe.Slot, pe.cI, pe.Token, pe.Line, pe.Column)
	exp := []string{}
	for _, e := range pe.Expected {
		exp = append(exp, e)
	}
	fmt.Fprintf(w, "Expected one of: [%s]", strings.Join(exp, ","))
	return w.String()
}

func (p *parser) parseError(slot slot.Label, i int, expected map[token.Type]string) {
	pe := &Error{cI: i, Slot: slot, Token: p.lex.Tokens[i], Expected: expected}
	p.parseErrors = append(p.parseErrors, pe)
}

func (p *parser) sortParseErrors() {
	sort.Slice(p.parseErrors,
		func(i, j int) bool {
			return p.parseErrors[j].Token.Lext() < p.parseErrors[i].Token.Lext()
		})
	for _, pe := range p.parseErrors {
		pe.Line, pe.Column = p.lex.GetLineColumn(pe.Token.Lext())
	}
}
