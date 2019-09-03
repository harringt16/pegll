# Gogll
**This is the Readme for gogll v0**

Gogll is a clustered nonterminal parser (CNP) generator following [Scott et al 2019](#Scott-et-al-2019). CNP is a version of generalised LL parsing (GLL). [GLL parsers can parse all context free (CF) languages](#Scott-et-al-2016).

Gogll generates **scannerless** parsers.

# Benefits
* General CF grammars are composable.
* [Operator precedence can be implemented very easily by disambiguating the parse forest](#Afroozeh-et-al-2019).

# Input Symbols, Markdown Files
Gogll accepts UTF-8 input strings. 
A gogll parser has two parse functions: 
* `Parse(I []byte) []*ParseError`
* `func ParseFile(fname string) []*ParseError`   
If `fname` ends with `.md` the parser ignores all text outside the markdown code blocks delimited by triple backticks. 

# Gogll Grammar
The [grammar for gogll v0](gogll.bnf) is a gocc BNF.


# Roadmap

* gogll v0: The first experimental working version of gogll generated with gocc.
* gogll v1: Planned. Generated by gogll v0. Code generation implemented from the gogll BSR.
* gogll v2: Planned. Generated gy gogll v1. 
  * Uses v1 code generation. 
  * Last vestiges of gocc BNF removed from gogll grammar.
  * Parser and BSR made objects -- global package variables removed

# Features considered for for future implementation
1. [EBNF](#Scott-et-al-2018)

# License
See [Apache license](License.txt)

# Notes
[see](notes.md)

# Changelog
[see](ChangeLog.md)

# Bibliography
<a name="Scott-et-al-2019"></a>
1. Elizabeth Scott, Adrian Johnstone and L. Thomas van Binsbergen.  
Derivation representation using binary subtree sets.  
In: Science of Computer Programming (175) 2019

<a name="Scott-et-al-2018"></a>
1. Elizabeth Scott and Adrian Johnstone.   
GLL Syntax Analysers For EBNF Grammars.   
In: [Science of Computer Programming
Volume 166, 15 November 2018](https://pure.royalholloway.ac.uk/portal/en/publications/gll-syntax-analysers-for-ebnf-grammars(58d1ec5e-28df-486a-879e-36d58a9f8abf).html)

<a name="Scott-et-al-2016"></a>
1. Elizabeth Scott and Adrian Johnstone.   
Structuring the GLL parsing algorithm for performance.   
In: [Science of Computer Programming
Volume 125, 1 September 2016](https://pure.royalholloway.ac.uk/portal/en/publications/structuring-the-gll-parsing-algorithm-for-performance(a95fc020-9918-4f17-a87a-845e2aee12b8).html)

<a name="Afroozeh-et-al-2013"></a>
1. Ali Afroozeh, Mark van den Brand, Adrian Johnstone, Elizabeth Scott, Jurgen Vinju.   
Safe Specification of Operator Precedence Rules.   
In: [Erwig M., Paige R.F., Van Wyk E. (eds) Software Language Engineering. SLE 2013. Lecture Notes in Computer Science, vol 8225. Springer, Cham](https://pure.royalholloway.ac.uk/portal/en/publications/safe-specification-of-operator-precedence-rules(0287d70e-92b8-4204-aafb-15a81de84968).html)