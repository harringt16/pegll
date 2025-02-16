# **`miniegg` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/miniegg.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
A modification of `miniegg` [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/miniegg.egg) parsing grammar ported into GoGLL to test an example structure.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`miniegg` GRAMMAR GUIDE**
The following grammar tests example structures to validate the `miniegg` grammar rules. 
```
package "miniegg"
```
`Grammar` is the semantic starting rule for the `miniegg` grammar. It calls for a space and then one or more repetitions of rules through `RepRule0x`.
```
Grammar         : " " Rule RepRule0x            ;
```
`Rule` is the semantic rule for a rule which is defined by an identifer followed by the character '=' while `RepRule0x` is the semantic rule for an expression repeating zero or more times until empty. This is accomplished by the ability for recursive calls in semantic rules and the ordered-choice `/` operator. See the [grammar for details.](../../gogll.md)
```
RepRule0x       : Rule RepRule0x
                / empty                         ;
Rule            : id eq " " Expr RepExpr0x      ;
```
`Expr` is the semantic rule for an expression which is defined by an identifer then any character other than '=' while `ExprRep` is the semantic rule for an expression repeating one or more times until empty. This is accomplished by the ability for recursive calls in semantic rules and the ordered-choice `/` operator. See the [grammar for details.](../../gogll.md)
```
RepExpr0x       : Expr RepExpr0x
                / empty                         ;
Expr            : id neq                        ; 
```
`id`, `eq`, and `neq` are lexical rules representing an identifier beginning with an uppercase letter followed by a space, the '=' character literal, and any character that is not '=' respectively. 
```
id              : upcase ' '                    ; 
eq              : '='                           ; 
neq             : not "="                       ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.