package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	newS := ""
	pNewS := ""
	for _, letter := range s {
		if ok, _ := regexp.MatchString("[a-zA-Z0-9]", string(letter)); !ok {
			continue
		}
		ll := strings.ToLower(string(letter))
		newS = newS + ll
		pNewS = ll + pNewS
	}
	return newS == pNewS
}
