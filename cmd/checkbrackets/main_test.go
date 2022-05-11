package main

import "testing"

func TestCheckBrackets(t *testing.T) {
	testingTable := []struct {
		brackets string
		expected bool
	}{
		{"[{}]", true},
		{"(()())", true},
		{"{]", false},
		{"[()]))()", false},
		{"[]{}({})", true},
	}

	for _, x := range testingTable {
		if checkBrackets(x.brackets) != x.expected {
			t.Errorf("Expected true, got false")
		}
	}
}
