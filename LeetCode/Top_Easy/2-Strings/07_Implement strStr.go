package main

import(
	"fmt"
)

func main(){


	input := "mississippi"
	input2 := "issip"
	fmt.Println(strStr(input,input2))


}


func strStr(haystack string, needle string) int {

	sizeH := len(haystack)
	sizeN := len(needle)

	if sizeH < sizeN{
		return -1
	}

	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == sizeN {
				return i
			}
			if i+j == sizeH {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
