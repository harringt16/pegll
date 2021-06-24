// Package parser is generated by gogll. Do not edit.
package parser

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"calc/lexer"
	"calc/parser/bsr"
	"calc/parser/slot"
	"calc/parser/symbols"
	"calc/token"
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
			{symbols.NT_EXPR, 0}: {},
		},
		crfNodes:    map[crfNode]*crfNode{},
		bsrSet:      bsr.New(symbols.NT_EXPR, l),
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
	p.ntAdd(symbols.NT_EXPR, 0)
	// p.DumpDescriptors()
	for !p.R.empty() {
		L, cU, p.cI = p.R.remove()

		// fmt.Println()
		// fmt.Printf("L:%s, cI:%d, I[p.cI]:%s, cU:%d\n", L, p.cI, p.lex.Tokens[p.cI], cU)
		// p.DumpDescriptors()

		switch L {
		case slot.CLOSE0R0: // CLOSE : ∙) space

			p.bsrSet.Add(slot.CLOSE0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.CLOSE0R1) {
				p.parseError(slot.CLOSE0R1, p.cI, first[slot.CLOSE0R1])
				break
			}

			p.bsrSet.Add(slot.CLOSE0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_CLOSE) {
				p.rtn(symbols.NT_CLOSE, cU, p.cI)
			} else {
				p.parseError(slot.CLOSE0R0, p.cI, followSets[symbols.NT_CLOSE])
			}
		case slot.DIVIDE0R0: // DIVIDE : ∙/ space

			p.bsrSet.Add(slot.DIVIDE0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.DIVIDE0R1) {
				p.parseError(slot.DIVIDE0R1, p.cI, first[slot.DIVIDE0R1])
				break
			}

			p.bsrSet.Add(slot.DIVIDE0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_DIVIDE) {
				p.rtn(symbols.NT_DIVIDE, cU, p.cI)
			} else {
				p.parseError(slot.DIVIDE0R0, p.cI, followSets[symbols.NT_DIVIDE])
			}
		case slot.ELEM0R0: // ELEM : ∙OPEN SUM CLOSE

			p.call(slot.ELEM0R1, cU, p.cI)
		case slot.ELEM0R1: // ELEM : OPEN ∙SUM CLOSE

			if !p.testSelect(slot.ELEM0R1) {
				p.parseError(slot.ELEM0R1, p.cI, first[slot.ELEM0R1])
				break
			}

			p.call(slot.ELEM0R2, cU, p.cI)
		case slot.ELEM0R2: // ELEM : OPEN SUM ∙CLOSE

			if !p.testSelect(slot.ELEM0R2) {
				p.parseError(slot.ELEM0R2, p.cI, first[slot.ELEM0R2])
				break
			}

			p.call(slot.ELEM0R3, cU, p.cI)
		case slot.ELEM0R3: // ELEM : OPEN SUM CLOSE ∙

			if p.follow(symbols.NT_ELEM) {
				p.rtn(symbols.NT_ELEM, cU, p.cI)
			} else {
				p.parseError(slot.ELEM0R0, p.cI, followSets[symbols.NT_ELEM])
			}
		case slot.ELEM1R0: // ELEM : ∙num

			p.bsrSet.Add(slot.ELEM1R1, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_ELEM) {
				p.rtn(symbols.NT_ELEM, cU, p.cI)
			} else {
				p.parseError(slot.ELEM1R0, p.cI, followSets[symbols.NT_ELEM])
			}
		case slot.EXPR0R0: // EXPR : ∙space SUM

			p.bsrSet.Add(slot.EXPR0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.EXPR0R1) {
				p.parseError(slot.EXPR0R1, p.cI, first[slot.EXPR0R1])
				break
			}

			p.call(slot.EXPR0R2, cU, p.cI)
		case slot.EXPR0R2: // EXPR : space SUM ∙

			if p.follow(symbols.NT_EXPR) {
				p.rtn(symbols.NT_EXPR, cU, p.cI)
			} else {
				p.parseError(slot.EXPR0R0, p.cI, followSets[symbols.NT_EXPR])
			}
		case slot.MINUS0R0: // MINUS : ∙- space

			p.bsrSet.Add(slot.MINUS0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.MINUS0R1) {
				p.parseError(slot.MINUS0R1, p.cI, first[slot.MINUS0R1])
				break
			}

			p.bsrSet.Add(slot.MINUS0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_MINUS) {
				p.rtn(symbols.NT_MINUS, cU, p.cI)
			} else {
				p.parseError(slot.MINUS0R0, p.cI, followSets[symbols.NT_MINUS])
			}
		case slot.OPEN0R0: // OPEN : ∙( space

			p.bsrSet.Add(slot.OPEN0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.OPEN0R1) {
				p.parseError(slot.OPEN0R1, p.cI, first[slot.OPEN0R1])
				break
			}

			p.bsrSet.Add(slot.OPEN0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_OPEN) {
				p.rtn(symbols.NT_OPEN, cU, p.cI)
			} else {
				p.parseError(slot.OPEN0R0, p.cI, followSets[symbols.NT_OPEN])
			}
		case slot.PLUS0R0: // PLUS : ∙+ space

			p.bsrSet.Add(slot.PLUS0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.PLUS0R1) {
				p.parseError(slot.PLUS0R1, p.cI, first[slot.PLUS0R1])
				break
			}

			p.bsrSet.Add(slot.PLUS0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_PLUS) {
				p.rtn(symbols.NT_PLUS, cU, p.cI)
			} else {
				p.parseError(slot.PLUS0R0, p.cI, followSets[symbols.NT_PLUS])
			}
		case slot.PLUSorMINUS0R0: // PLUSorMINUS : ∙PLUS PROD

			p.call(slot.PLUSorMINUS0R1, cU, p.cI)
		case slot.PLUSorMINUS0R1: // PLUSorMINUS : PLUS ∙PROD

			if !p.testSelect(slot.PLUSorMINUS0R1) {
				p.parseError(slot.PLUSorMINUS0R1, p.cI, first[slot.PLUSorMINUS0R1])
				break
			}

			p.call(slot.PLUSorMINUS0R2, cU, p.cI)
		case slot.PLUSorMINUS0R2: // PLUSorMINUS : PLUS PROD ∙

			if p.follow(symbols.NT_PLUSorMINUS) {
				p.rtn(symbols.NT_PLUSorMINUS, cU, p.cI)
			} else {
				p.parseError(slot.PLUSorMINUS0R0, p.cI, followSets[symbols.NT_PLUSorMINUS])
			}
		case slot.PLUSorMINUS1R0: // PLUSorMINUS : ∙MINUS PROD

			p.call(slot.PLUSorMINUS1R1, cU, p.cI)
		case slot.PLUSorMINUS1R1: // PLUSorMINUS : MINUS ∙PROD

			if !p.testSelect(slot.PLUSorMINUS1R1) {
				p.parseError(slot.PLUSorMINUS1R1, p.cI, first[slot.PLUSorMINUS1R1])
				break
			}

			p.call(slot.PLUSorMINUS1R2, cU, p.cI)
		case slot.PLUSorMINUS1R2: // PLUSorMINUS : MINUS PROD ∙

			if p.follow(symbols.NT_PLUSorMINUS) {
				p.rtn(symbols.NT_PLUSorMINUS, cU, p.cI)
			} else {
				p.parseError(slot.PLUSorMINUS1R0, p.cI, followSets[symbols.NT_PLUSorMINUS])
			}
		case slot.PROD0R0: // PROD : ∙ELEM ToDRep

			p.call(slot.PROD0R1, cU, p.cI)
		case slot.PROD0R1: // PROD : ELEM ∙ToDRep

			if !p.testSelect(slot.PROD0R1) {
				p.parseError(slot.PROD0R1, p.cI, first[slot.PROD0R1])
				break
			}

			p.call(slot.PROD0R2, cU, p.cI)
		case slot.PROD0R2: // PROD : ELEM ToDRep ∙

			if p.follow(symbols.NT_PROD) {
				p.rtn(symbols.NT_PROD, cU, p.cI)
			} else {
				p.parseError(slot.PROD0R0, p.cI, followSets[symbols.NT_PROD])
			}
		case slot.PoMRep0R0: // PoMRep : ∙PLUSorMINUS PoMRep

			p.call(slot.PoMRep0R1, cU, p.cI)
		case slot.PoMRep0R1: // PoMRep : PLUSorMINUS ∙PoMRep

			if !p.testSelect(slot.PoMRep0R1) {
				p.parseError(slot.PoMRep0R1, p.cI, first[slot.PoMRep0R1])
				break
			}

			p.call(slot.PoMRep0R2, cU, p.cI)
		case slot.PoMRep0R2: // PoMRep : PLUSorMINUS PoMRep ∙

			if p.follow(symbols.NT_PoMRep) {
				p.rtn(symbols.NT_PoMRep, cU, p.cI)
			} else {
				p.parseError(slot.PoMRep0R0, p.cI, followSets[symbols.NT_PoMRep])
			}
		case slot.PoMRep1R0: // PoMRep : ∙
			p.bsrSet.AddEmpty(slot.PoMRep1R0, p.cI)

			if p.follow(symbols.NT_PoMRep) {
				p.rtn(symbols.NT_PoMRep, cU, p.cI)
			} else {
				p.parseError(slot.PoMRep1R0, p.cI, followSets[symbols.NT_PoMRep])
			}
		case slot.SUM0R0: // SUM : ∙PROD PoMRep

			p.call(slot.SUM0R1, cU, p.cI)
		case slot.SUM0R1: // SUM : PROD ∙PoMRep

			if !p.testSelect(slot.SUM0R1) {
				p.parseError(slot.SUM0R1, p.cI, first[slot.SUM0R1])
				break
			}

			p.call(slot.SUM0R2, cU, p.cI)
		case slot.SUM0R2: // SUM : PROD PoMRep ∙

			if p.follow(symbols.NT_SUM) {
				p.rtn(symbols.NT_SUM, cU, p.cI)
			} else {
				p.parseError(slot.SUM0R0, p.cI, followSets[symbols.NT_SUM])
			}
		case slot.TIMES0R0: // TIMES : ∙* space

			p.bsrSet.Add(slot.TIMES0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.TIMES0R1) {
				p.parseError(slot.TIMES0R1, p.cI, first[slot.TIMES0R1])
				break
			}

			p.bsrSet.Add(slot.TIMES0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_TIMES) {
				p.rtn(symbols.NT_TIMES, cU, p.cI)
			} else {
				p.parseError(slot.TIMES0R0, p.cI, followSets[symbols.NT_TIMES])
			}
		case slot.TIMESorDIVIDE0R0: // TIMESorDIVIDE : ∙TIMES ELEM

			p.call(slot.TIMESorDIVIDE0R1, cU, p.cI)
		case slot.TIMESorDIVIDE0R1: // TIMESorDIVIDE : TIMES ∙ELEM

			if !p.testSelect(slot.TIMESorDIVIDE0R1) {
				p.parseError(slot.TIMESorDIVIDE0R1, p.cI, first[slot.TIMESorDIVIDE0R1])
				break
			}

			p.call(slot.TIMESorDIVIDE0R2, cU, p.cI)
		case slot.TIMESorDIVIDE0R2: // TIMESorDIVIDE : TIMES ELEM ∙

			if p.follow(symbols.NT_TIMESorDIVIDE) {
				p.rtn(symbols.NT_TIMESorDIVIDE, cU, p.cI)
			} else {
				p.parseError(slot.TIMESorDIVIDE0R0, p.cI, followSets[symbols.NT_TIMESorDIVIDE])
			}
		case slot.TIMESorDIVIDE1R0: // TIMESorDIVIDE : ∙DIVIDE ELEM

			p.call(slot.TIMESorDIVIDE1R1, cU, p.cI)
		case slot.TIMESorDIVIDE1R1: // TIMESorDIVIDE : DIVIDE ∙ELEM

			if !p.testSelect(slot.TIMESorDIVIDE1R1) {
				p.parseError(slot.TIMESorDIVIDE1R1, p.cI, first[slot.TIMESorDIVIDE1R1])
				break
			}

			p.call(slot.TIMESorDIVIDE1R2, cU, p.cI)
		case slot.TIMESorDIVIDE1R2: // TIMESorDIVIDE : DIVIDE ELEM ∙

			if p.follow(symbols.NT_TIMESorDIVIDE) {
				p.rtn(symbols.NT_TIMESorDIVIDE, cU, p.cI)
			} else {
				p.parseError(slot.TIMESorDIVIDE1R0, p.cI, followSets[symbols.NT_TIMESorDIVIDE])
			}
		case slot.ToDRep0R0: // ToDRep : ∙TIMESorDIVIDE ToDRep

			p.call(slot.ToDRep0R1, cU, p.cI)
		case slot.ToDRep0R1: // ToDRep : TIMESorDIVIDE ∙ToDRep

			if !p.testSelect(slot.ToDRep0R1) {
				p.parseError(slot.ToDRep0R1, p.cI, first[slot.ToDRep0R1])
				break
			}

			p.call(slot.ToDRep0R2, cU, p.cI)
		case slot.ToDRep0R2: // ToDRep : TIMESorDIVIDE ToDRep ∙

			if p.follow(symbols.NT_ToDRep) {
				p.rtn(symbols.NT_ToDRep, cU, p.cI)
			} else {
				p.parseError(slot.ToDRep0R0, p.cI, followSets[symbols.NT_ToDRep])
			}
		case slot.ToDRep1R0: // ToDRep : ∙
			p.bsrSet.AddEmpty(slot.ToDRep1R0, p.cI)

			if p.follow(symbols.NT_ToDRep) {
				p.rtn(symbols.NT_ToDRep, cU, p.cI)
			} else {
				p.parseError(slot.ToDRep1R0, p.cI, followSets[symbols.NT_ToDRep])
			}

		default:
			panic("This must not happen")
		}
	}
	if !p.bsrSet.Contain(symbols.NT_EXPR, 0, m) {
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
	_, exist := first[l][p.lex.Tokens[p.cI].Type()]
	// fmt.Printf("testSelect(%s) = %t\n", l, exist)
	return exist
}

var first = []map[token.Type]string{
	// CLOSE : ∙) space
	{
		token.T_1: ")",
	},
	// CLOSE : ) ∙space
	{
		token.T_7: "space",
	},
	// CLOSE : ) space ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// DIVIDE : ∙/ space
	{
		token.T_5: "/",
	},
	// DIVIDE : / ∙space
	{
		token.T_7: "space",
	},
	// DIVIDE : / space ∙
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// ELEM : ∙OPEN SUM CLOSE
	{
		token.T_0: "(",
	},
	// ELEM : OPEN ∙SUM CLOSE
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// ELEM : OPEN SUM ∙CLOSE
	{
		token.T_1: ")",
	},
	// ELEM : OPEN SUM CLOSE ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// ELEM : ∙num
	{
		token.T_6: "num",
	},
	// ELEM : num ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// EXPR : ∙space SUM
	{
		token.T_7: "space",
	},
	// EXPR : space ∙SUM
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// EXPR : space SUM ∙
	{
		token.EOF: "$",
	},
	// MINUS : ∙- space
	{
		token.T_4: "-",
	},
	// MINUS : - ∙space
	{
		token.T_7: "space",
	},
	// MINUS : - space ∙
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// OPEN : ∙( space
	{
		token.T_0: "(",
	},
	// OPEN : ( ∙space
	{
		token.T_7: "space",
	},
	// OPEN : ( space ∙
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUS : ∙+ space
	{
		token.T_3: "+",
	},
	// PLUS : + ∙space
	{
		token.T_7: "space",
	},
	// PLUS : + space ∙
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUSorMINUS : ∙PLUS PROD
	{
		token.T_3: "+",
	},
	// PLUSorMINUS : PLUS ∙PROD
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUSorMINUS : PLUS PROD ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PLUSorMINUS : ∙MINUS PROD
	{
		token.T_4: "-",
	},
	// PLUSorMINUS : MINUS ∙PROD
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUSorMINUS : MINUS PROD ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PROD : ∙ELEM ToDRep
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PROD : ELEM ∙ToDRep
	{
		token.T_2: "*",
		token.T_5: "/",
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PROD : ELEM ToDRep ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PoMRep : ∙PLUSorMINUS PoMRep
	{
		token.T_3: "+",
		token.T_4: "-",
	},
	// PoMRep : PLUSorMINUS ∙PoMRep
	{
		token.T_3: "+",
		token.T_4: "-",
		token.EOF: "$",
		token.T_1: ")",
	},
	// PoMRep : PLUSorMINUS PoMRep ∙
	{
		token.EOF: "$",
		token.T_1: ")",
	},
	// PoMRep : ∙
	{
		token.EOF: "$",
		token.T_1: ")",
	},
	// SUM : ∙PROD PoMRep
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// SUM : PROD ∙PoMRep
	{
		token.T_3: "+",
		token.T_4: "-",
		token.EOF: "$",
		token.T_1: ")",
	},
	// SUM : PROD PoMRep ∙
	{
		token.EOF: "$",
		token.T_1: ")",
	},
	// TIMES : ∙* space
	{
		token.T_2: "*",
	},
	// TIMES : * ∙space
	{
		token.T_7: "space",
	},
	// TIMES : * space ∙
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// TIMESorDIVIDE : ∙TIMES ELEM
	{
		token.T_2: "*",
	},
	// TIMESorDIVIDE : TIMES ∙ELEM
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// TIMESorDIVIDE : TIMES ELEM ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// TIMESorDIVIDE : ∙DIVIDE ELEM
	{
		token.T_5: "/",
	},
	// TIMESorDIVIDE : DIVIDE ∙ELEM
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// TIMESorDIVIDE : DIVIDE ELEM ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// ToDRep : ∙TIMESorDIVIDE ToDRep
	{
		token.T_2: "*",
		token.T_5: "/",
	},
	// ToDRep : TIMESorDIVIDE ∙ToDRep
	{
		token.T_2: "*",
		token.T_5: "/",
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// ToDRep : TIMESorDIVIDE ToDRep ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// ToDRep : ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
}

var followSets = []map[token.Type]string{
	// CLOSE
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// DIVIDE
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// ELEM
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// EXPR
	{
		token.EOF: "$",
	},
	// MINUS
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// OPEN
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUS
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// PLUSorMINUS
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PROD
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
	},
	// PoMRep
	{
		token.EOF: "$",
		token.T_1: ")",
	},
	// SUM
	{
		token.EOF: "$",
		token.T_1: ")",
	},
	// TIMES
	{
		token.T_0: "(",
		token.T_6: "num",
	},
	// TIMESorDIVIDE
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: "*",
		token.T_3: "+",
		token.T_4: "-",
		token.T_5: "/",
	},
	// ToDRep
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_3: "+",
		token.T_4: "-",
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
