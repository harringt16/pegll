//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
package ast

import (
	"fmt"
	"gogll/goutil/stringset"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

// symbol.Value() -> Symbols
var symbols = make(map[string]Symbol)

// var nonTerminals = make(map[string]bool)
var rules = make(map[string]*Rule)

// token literal -> token type
var terminals = make(map[string]string)
var startSymbol = ""
var parserPackage = ""

// func addTerminal(sym Symbol) {
// 	terminals[sym.Value()] = true
// }

func addRule(rule *Rule) error {
	if _, exist := rules[rule.Head.stringValue]; exist {
		return fmt.Errorf("Duplicate declaration of rule %s", rule.Head.stringValue)
	}
	rules[rule.Head.stringValue] = rule
	AddSymbol(rule.Head)
	if rule.IsStartSymbol {
		if startSymbol != "" {
			return fmt.Errorf("Duplicate start symbol %s", rule.Head.stringValue)
		}
		startSymbol = rule.Head.stringValue
	}
	return nil
}

func AddSymbol(s Symbol) (Symbol, error) {
	// fmt.Printf("symtab.AddSymbol(\"%s\")\n", s.Value())
	if str, ok := s.(*String); ok {
		if err := addStringSymbols(str); err != nil {
			return nil, err
		}
		return s, nil
	}
	if s1, exist := symbols[s.StringValue()]; exist && !s1.Equal(s) {
		return nil, fmt.Errorf("Incompatible duplicate symbol %s", s.StringValue())
	}
	symbols[s.StringValue()] = s
	return s, nil
}

func addStringSymbols(s *String) error {
	syms, err := newStringChars(s)
	if err != nil {
		return err
	}
	for _, sym := range syms {
		if _, err := AddSymbol(sym); err != nil {
			return err
		}
	}
	return nil
}

func GetNonTerminals() (ss Symbols) {
	for nt, sym := range symbols {
		if !sym.IsTerminal() {
			ss = append(ss, nt)
		}
	}
	sort.Strings(ss)
	return
}

func GetPackage() string {
	return parserPackage
}

func GetRule(head string) *Rule {
	if r, exist := rules[head]; !exist {
		panic(fmt.Sprintf("No rule %s", head))
	} else {
		return r
	}
}

func GetRules() map[string]*Rule {
	return rules
}

func GetStartSymbol() string {
	if startSymbol == "" {
		fmt.Printf("Error: No Start Symbol specified in the grammar\n")
		os.Exit(1)
	}
	return startSymbol
}

func GetTerminals() (ss Symbols) {
	for t, sym := range symbols {
		if sym.IsTerminal() {
			ss = append(ss, t)
		}
	}
	sort.Strings(ss)
	return
}

func GetSymbol(symbolName string) Symbol {
	if s, exist := symbols[symbolName]; exist {
		return s
	}
	panic(fmt.Sprintf("No symbol %s", symbolName))
}

func GetSymbols() Symbols {
	ss := stringset.New()
	for _, sym := range symbols {
		ss.Add([]string(sym.Symbols())...)
	}
	ss1 := ss.Elements()
	sort.Strings(ss1)
	return ss1
}

func IsTerminal(symbol string) bool {
	sym, exist := symbols[symbol]
	if !exist {
		panic(fmt.Sprintf("Symbol %s does not exist", symbol))
	}
	return sym.IsTerminal()
}

type Symbols []string

func (ss Symbols) Contain(sym string) bool {
	for _, s := range ss {
		if s == sym {
			return true
		}
	}
	return false
}

func (ss Symbols) Remove(s string) Symbols {
	newSS := Symbols{}
	for _, s1 := range ss {
		if s1 != s {
			newSS = append(newSS, s1)
		}
	}
	return newSS
}

func (ss Symbols) RemoveDuplicates() (newSS Symbols) {
	seen := make(map[string]bool)
	for _, s := range ss {
		if _, exist := seen[s]; !exist {
			newSS = append(newSS, s)
			seen[s] = true
		}
	}
	return
}

// func (ss Symbols) Len() int {
// 	return len(ss)
// }

// func (ss Symbols) Less(i, j int) bool {
// 	return ss[i] < ss[j]
// }

// func (ss Symbols) Swap(i, j int) {
// 	iTmp := ss[i]
// 	ss[i] = ss[j]
// 	ss[j] = iTmp
// }

/*** Utils ***/

func isNT(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

/*** Dump ***/

func DumpSymbols() {
	for _, s := range GetSymbols() {
		fmt.Println(s)
	}
}

func DumpNT() {
	fmt.Println("NT: Alternative")
	for _, id := range GetNonTerminals() {
		DumpRuleSymbols(id)
	}
	fmt.Println()
}

func DumpT() {
	fmt.Println("Terminals")
	for _, nt := range GetTerminals() {
		fmt.Println("  ", nt)
	}
	fmt.Println()
}

func DumpRuleSymbols(id string) {
	rule := GetRule(id)
	for _, a := range rule.Alternates {
		fmt.Printf("  %s: %s\n", id, a.Symbols())
	}
}
