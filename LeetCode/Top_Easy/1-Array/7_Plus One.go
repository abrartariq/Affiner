package main

import "fmt"

func main(){
	Input := []int{1,1,3,5,6,3,5}
	fmt.Println(plusOne(Input))
	
}


func plusOne(digits []int) []int {

    newDigits := []int{0}

	newDigits = append(newDigits, digits...)
	
    for i := len(newDigits) - 1; i >= 0; i-- {
    	if newDigits[i] == 9{
    		newDigits[i] = 0
    	}else{
    		newDigits[i]++
    		break
    	}
    }


    if newDigits[0] == 1{
    	return newDigits
    }

    return newDigits[1:]
}