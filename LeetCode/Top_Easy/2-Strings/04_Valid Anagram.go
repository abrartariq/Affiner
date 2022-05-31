package main

import (
    "fmt"
)

func main(){
	fmt.Println(isAnagram("92313","12393"))
}

func isAnagram(s string, t string) bool {
    wordMap := make(map[byte]int)

    stringSize := len(s)

    if len(s) != len(t){
    	return false
    }

    for i:=0; i<stringSize; i++{

        val, isPresent := wordMap[s[i]]
        if isPresent{
            wordMap[s[i]] = val + 1
        }else{
            wordMap[s[i]] = 1 
        }
    }

	for i:=0; i<stringSize; i++{

        val, isPresent := wordMap[t[i]]
        if isPresent && val > 0{
            wordMap[t[i]] = val - 1
        }else{
             return false
        }
    }
    
    return true
}
