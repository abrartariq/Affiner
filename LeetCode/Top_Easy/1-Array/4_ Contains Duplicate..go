package main

import "fmt"

func main(){

	Input := []int{1,2,3,4,5,6,7}
	fmt.Println(containsDuplicate(Input))
	// rotate(Input,3)


}

func containsDuplicate(nums []int) bool {
    
    myMap := make(map[int]int)

    for i := 0; i < len(nums); i++{

    	_, found := myMap[nums[i]]

    	if found == true{
    		return true
    	}else{
    		myMap[nums[i]] = 1
    	}
    } 

    return false
}