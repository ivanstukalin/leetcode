package main

//O(n)
func isValid(s string) bool {
	var stack []string
	bracketMap := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for i:=0; i<len(s); i++ {
		bracket := s[i:i+1]
		if bracket == "}" || bracket == "]" || bracket == ")" {
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[bracket] {
				return false;
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, bracket)
		}
	}
	return len(stack) == 0;
}
