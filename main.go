/*
Given two strings, write a method to decide if one is a permutation
of the other.
*/

package main

import (
	"fmt"
)

var a string
var b string
var c string

func isPalindrome(a string, b string) bool {
	//check if they are same length
	if len(a) > len(b) || len(b) > len(a) {
		return false
	} else {
		if a == reverseString(b) {
			return true
		} else {
			return false
		}

	}
}

func reverseString(a string) string {
	r := []rune(a)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	a = "god"
	b = "dog"
	c = "red"

	fmt.Println(isPalindrome(a, b))
	fmt.Println(isPalindrome(b, c))

}
