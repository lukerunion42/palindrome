package main

import "testing"

type PalindromeTable struct {
	input    string
	input2   string
	expected bool
}

var palindromeResults = []PalindromeTable{
	{"god", "dog", true},
	{"racecar", "racecar", true},
	{"red", "bus", false},
	{"seventeen", "neetneves", true},
}

func TestPalindromeTable(t *testing.T) {
	for _, test := range palindromeResults {
		result := isPalindrome(test.input, test.input2)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}

}
