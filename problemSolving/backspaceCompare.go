package main

import (
	"fmt"
	"strings"
)

type stack []string

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(value rune) {
	*s = append(*s, string(value))
}

func (s *stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func actualStr(s string) string {
	st := stack{}
	for _, r := range s {
		st.Push(r)
		if r == '#' {
			_ = st.Pop()
			_ = st.Pop()
		}
	}
	return strings.Join(st, "")
}

func backspaceCompare(s string, t string) bool {
	valS := actualStr(s)
	valT := actualStr(t)
	return valS == valT
}

/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.


*/
func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(backspaceCompare(s, t))
}
