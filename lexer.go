package main

import (
	"fmt"
)

type TokenClass int

const (
	Null TokenClass = iota
	Number
	String
	Identifier

	// Keywords
	Breakpoint1 // This is used to make the following start at 401, 402 .. dynamically
	Print = iota + (100 * Breakpoint1) - Breakpoint1
	If
	Goto
	Input
	Let
	Gosub
	Return
	Clear
	List
	Run
	End

	// Operators
	Breakpoint2 = iota
	GT = iota + (100 * (Breakpoint1 + 1)) - Breakpoint2
	GTOE
	LT
	LTOE
	EQ 
	EQEQ 
	NEQ
	Plus
	Minus
	Times
	Divide
)

type lexedToken struct {
	text string
	class TokenClass 
}

type lexedLine []lexedToken

func lex(in string) (o []lexedLine, err error){
	o = make([]lexedLine, 0)
	in = fmt.Sprintf("%s\n",in) // To make parsing the last line consistent
	ls := 0;
	ts := 0;
	ta := make([]string, 0);
	for te := 0; te < len(in); te++{
		if in[te] == '\n' { // Line ended, lex the current line and move on
			if te > ts { // If the token is just a space, ignore
				ta = append(ta, in[ts:te])
			}
			if len(ta) > 0 {
				l, e := lexLine(ta);
				if (e != nil) {
					panic(e)
				}
				if len(l) > 0 {
					o = append(o, l)
				}
				ta = make([]string, 0)
			}
			ls = te + 1;
			ts = ls;
		} else if in[te] == ' ' && in[ts] != '"' { // Token ended, append to ta. Spaces are only allowed inside a string
			if te > ts { // If the token is just a space, ignore
				ta = append(ta, in[ts:te])
			}
			ts = te + 1
		} else if in[te] == '"' && in[ts] == '"' && te > ts {
			ta = append(ta, in[ts:te + 1])
			ts = te + 1
		}
	}
	return
}

func lexLine(ta []string) (o lexedLine, err error){
	if ta[0][0:2] == "//" { // Skip comments
		return
	}
	o = make([]lexedToken, len(ta));
	for i:= range(ta) {
		t := ta[i]
		to := parseKeyword(t)
		if to.class == Null {
			to = parseOperator(t)
		}
		if to.class == Null {
			to = parseNumber(t)
		}
		if to.class == Null {
			to = parseString(t)
		}
		if to.class == Null {
			to = parseIdentifier(t)
		}
		if to.class != Null {
			o[i] = to
		} else {
			err = fmt.Errorf("unparseable token %s", t)
		}
	}
	return
}

func parseKeyword(t string) (to lexedToken) {
	if t == "PRINT" {
		to = lexedToken {
			t,
			Print,
		}
	} else if t == 	"IF" {
		to = lexedToken {
			t,
			If,
		}
	} else if t == "GOTO" {
		to = lexedToken {
			t,
			Goto,
		}
	} else if t == "INPUT" {
		to = lexedToken {
			t,
			Input,
		}
	} else if t == "LET" {
		to = lexedToken {
			t,
			Let,
		}
	} else if t == "GOSUB" {
		to = lexedToken {
			t,
			Gosub,
		}
	} else if t == "RETURN" {
		to = lexedToken {
			t,
			Return,
		}
	} else if t == "CLEAR" {
		to = lexedToken {
			t,
			Clear,
		}
	} else if t == "LIST" {
		to = lexedToken {
			t,
			List,
		}
	} else if t == "RUN" {
		to = lexedToken {
			t,
			Run,
		}
	} else if t == "END" {
		to = lexedToken {
			t,
			End,
		}
	}
	return
}

func parseOperator(t string) (to lexedToken) {
		if t == ">" {
			to = lexedToken{
				t,
				GT,
			}
		} else if t == ">=" {
			to = lexedToken{
				t,
				GTOE,
			}
		} else if t == "<" {
			to = lexedToken{
				t,
				LT,
			}
		} else if t == "<=" {
			to = lexedToken{
				t,
				LTOE,
			}
		} else if t == "=" {
			to = lexedToken{
				t,
				EQ,
			}
		} else if t == "==" {
			to = lexedToken{
				t,
				EQEQ,
			}
		} else if t == "!=" {
			to = lexedToken{
				t,
				NEQ,
			}
		} else if t == "+" {
			to = lexedToken{
				t,
				Plus,
			}
		} else if t == "-" {
			to = lexedToken{
				t,
				Minus,
			}
		} else if t == "*" {
			to = lexedToken{
				t,
				Times,
			}
		} else if t == "/" {
			to = lexedToken{
				t,
				Divide,
			}
		}
		return
}

func parseNumber(t string) (to lexedToken) {
	result := true
	dc := 0
	for i := range(t) {
		if t[i] < '0' || t[i] > '9' {
			if t[i] == '.' {
				dc += 1;
				if i == 0 || i == len(t) - 1 || dc > 1 { // A decimal point is only allowed once in a number, and never at the very beginning or end.
					result = false
					break
				}
			} else {
				result = false
				break
			}
		}
	}
	if (result) {
		to = lexedToken{ 
			t,
			Number,
		}
	}
	return
}

func parseString(t string) (to lexedToken) {
	if t[0] == t[len(t) - 1] && t[0] == '"' {
		to = lexedToken{
			t,
			String,
		}
	}
	return
}

func parseIdentifier(t string) (to lexedToken) {
	for i := range(t) {
		if t[i] > 'Z' || t[i] < 'A' {
			return
		} 
	}
	to = lexedToken{ 
		t,
		Identifier,
	}
	return
}