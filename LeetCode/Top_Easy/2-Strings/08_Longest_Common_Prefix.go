package main

import(
	"fmt"
)

func main(){


	input := []string{"mississippi", "miss","missis","missip"}
	fmt.Println(longestCommonPrefix(input))


}


func longestCommonPrefix(strs []string) string {
    
    minString := strs[0]
	minSize := len(minString)

	for i := 0; i < len(strs); i++ {
		cSize := len(strs[i])
		if cSize < minSize{
			minString = strs[i]
			minSize = cSize
		}
	}
	
	for i := 0; i < minSize; i++ {
		for j := 0; j < len(strs); j++{
			if strs[j][i] != minString[i]{
				return minString[:i]
			}
		}
	}

	return minString
}