
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"axbc/parser/symbols"
)

type Label int

const(
	AxBC0R0 Label = iota
	AxBC0R1
	AxBC0R2
	AxBC1R0
	AxBC1R1
	AxBC1R2
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	AxBC0R0: {
		symbols.NT_AxBC, 0, 0, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_2,
		}, 
		AxBC0R0, 
	},
	AxBC0R1: {
		symbols.NT_AxBC, 0, 1, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_2,
		}, 
		AxBC0R1, 
	},
	AxBC0R2: {
		symbols.NT_AxBC, 0, 2, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_2,
		}, 
		AxBC0R2, 
	},
	AxBC1R0: {
		symbols.NT_AxBC, 1, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_2,
		}, 
		AxBC1R0, 
	},
	AxBC1R1: {
		symbols.NT_AxBC, 1, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_2,
		}, 
		AxBC1R1, 
	},
	AxBC1R2: {
		symbols.NT_AxBC, 1, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_2,
		}, 
		AxBC1R2, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_AxBC,0,0 }: AxBC0R0,
	Index{ symbols.NT_AxBC,0,1 }: AxBC0R1,
	Index{ symbols.NT_AxBC,0,2 }: AxBC0R2,
	Index{ symbols.NT_AxBC,1,0 }: AxBC1R0,
	Index{ symbols.NT_AxBC,1,1 }: AxBC1R1,
	Index{ symbols.NT_AxBC,1,2 }: AxBC1R2,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_AxBC:[]Label{ AxBC0R0,AxBC1R0 },
}

