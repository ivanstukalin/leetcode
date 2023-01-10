package main

func firstUniqChar(s string) int {
	runes := map[rune]int{}
	nonUniques := map[rune]int{}
	for pos, letter := range s {
		_, ok := runes[letter]
		if ok {
			nonUniques[letter] = pos
			continue
		}
		runes[letter] = pos
	}

	for pos, letter := range s {
		_, ok := nonUniques[letter]
		if !ok {
			return pos
		}
	}
	return -1
}
