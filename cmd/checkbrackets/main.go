package main

import "github.com/vitor-mariano/data-structures-and-algorithms/pkg/stack"

func checkBrackets(brackets string) bool {
	s := stack.New[rune]()
	pairs := []struct {
		open  rune
		close rune
	}{
		{'{', '}'},
		{'[', ']'},
		{'(', ')'},
	}

	for _, char := range brackets {
		for _, pair := range pairs {
			switch {
			case char == pair.open:
				s.Push(char)
			case char == pair.close:
				if s.GetSize() == 0 || s.Pop() != pair.open {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	s := []string{"[{}]", "(()())", "{]", "[()]))()", "[]{}({})"}

	for _, x := range s {
		if checkBrackets(x) {
			println(x, "is valid")
		} else {
			println(x, "is not valid")
		}
	}
}
