package main

func ReverseString(s []byte)  {
    last := len(s)-1
    for key, value := range s {
        if key > last-key-1 {
            break;
        }
        s[key] = s[last-key]
        s[last-key] = value
    } 
}