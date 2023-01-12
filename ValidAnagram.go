package main

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    lettersS := map[byte]int{}
    lettersT := map[byte]int{}
    
    for i := 0; i < len(s); i++ {
        _, okS := lettersS[s[i]]
        if okS {
            lettersS[s[i]]++
        } else {
            lettersS[s[i]] = 1
        }
        
        _, okT := lettersT[t[i]]
        if okT {
            lettersT[t[i]]++
        } else {
            lettersT[t[i]] = 1
        }
    }
    
    for letter, amountS := range lettersS {
        amountT, ok := lettersT[letter]
        if !ok || amountT != amountS {
            return false
        }
    }
    
    return true
}