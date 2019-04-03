/*
Given two strings, write a method to decide if one is a permutation
of the other.
*/

package main

import "testing"

type PermutationTable struct {
	input    string
	input2   string
	expected bool
}

var permutationResults = []PermutationTable{
	{"god", "dog", true},
	{"racecar", "racecar", true},
	{"red", "bus", false},
	{"seventeen", "neetneves", true},
}

func TestPermutationTable(t *testing.T) {
	for _, test := range permutationResults {
		result := isPermutation(test.input, test.input2)
		if result != test.expected {
			t.Fatal("Expected Result Not Given")
		}
	}

}
