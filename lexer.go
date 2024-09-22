package main

import (
	"errors"
	"fmt"
)

type TokenClass int

const (
	KeyWord TokenClass = iota
	Operator
	Number
	String
	Identifier
)

type lexed_token struct {
	text string
	class TokenClass 
}

type lexed_line []lexed_token

func lex(in string) (o []lexed_line, err error){
	o = make([]lexed_line, 0)
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
				l, e := lex_line(ta);
				if ( e == nil) {
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

func lex_line(ta []string) (o lexed_line, err error){
	if ta[0][0:2] == "//" { // Skip comments
		err = errors.New("comment")
		return
	}

	// for i:= range(ta) {

	// }
	return
}
