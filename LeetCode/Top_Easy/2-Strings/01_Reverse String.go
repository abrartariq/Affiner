package main

import "fmt"

func main(){
	Input := []byte("ABC")
	reverseString(Input)
	fmt.Println(Input)
}


func reverseString(s []byte)  {
    sizeA := len(s)
    for i:=0;i<sizeA/2;i++{
        s[i], s[sizeA-1-i] = s[sizeA-1-i], s[i] 
    }
}