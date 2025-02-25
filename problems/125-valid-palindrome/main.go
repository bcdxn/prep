package main

import (
	"fmt"
	"strings"
)

func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println("result: ", res)
}

const (
	zero = byte('0')
	nine = byte('9')
	A    = byte('A')
	Z    = byte('Z')
	a    = byte('a')
	z    = byte('z')
)

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for !isAlpha(s[l]) && l < r {
			l++
		}
		for !isAlpha(s[r]) && l < r {
			r--
		}

		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlpha(str byte) bool {
	return (str >= zero && str <= nine) ||
		(str >= A && str <= Z) ||
		(str >= a && str <= z)
}
