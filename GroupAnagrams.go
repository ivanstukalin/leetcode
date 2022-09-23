package main

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	var charSet = map[int]int{}
	for key, str := range strs {
		if str == "" {
			charSet[key] = 0
			continue
		}
		for _, letter := range []byte(str) {
			charSet[key] += int(letter) * int(letter) * int(letter) * int(letter)
		}
	}

	for str, bytes := range charSet {
		temp := []string{strs[str]}
		for key, bytes2 := range charSet {
			if str != key && bytes == bytes2 {
				temp = append(temp, strs[key])
				delete(charSet, key)
			}
		}
		result = append(result, temp)
	}
	return result
}
