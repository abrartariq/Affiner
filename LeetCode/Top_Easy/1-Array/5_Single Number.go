package main

import "fmt"

func main(){

	Input := []int{1,1,3,3,5,6,5}
	fmt.Println(singleNumber(Input))
	// rotate(Input,3)


}

func singleNumber(nums []int) int {
    
    ans := nums[0]

    for i := 1 ; i < len(nums); i++{
    	ans = ans ^ nums[i]
    } 

    return ans
}