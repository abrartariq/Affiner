package main

import (
    "fmt"
)

func main(){
	fmt.Println(firstUniqChar("9231"))
}

    
func firstUniqChar(s string) int {
    seenMap := make([]int, 128)
    
    for i := len(s) - 1; i >= 0; i--{
        seenMap[byte(s[i])]++
    }
    for i := 0; i < len(s); i++ {
        if seenMap[byte(s[i])] == 1 {
            return i
        }
    }
    return -1
}
