package main

import "fmt"

func main(){


	Input := []int{2,7,11,15}
	fmt.Println(twoSum(Input,9))

}


func twoSum(nums []int, target int) []int {
    
	sizeA := len(nums)
	myMap := make(map[int]int)
	ansVal := []int{-1,-1}

	for i := 0; i < sizeA; i++ {
		iVal, isPresent := myMap[target - nums[i]]
		if !isPresent {
			myMap[nums[i]] = i
		}else{
			ansVal[0] = i
			ansVal[1] = iVal
			break
		}
	}

	return ansVal

}