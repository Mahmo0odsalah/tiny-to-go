package main

import "fmt"

func main() {
	o, _ := lex("LET 13 = \"blom s\"\n LET  AC = 12\n")
	for i:= range(o) {
		fmt.Println(o[i])
	}
}
