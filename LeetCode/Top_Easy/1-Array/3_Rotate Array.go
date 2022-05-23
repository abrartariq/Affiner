package main

import "fmt"

func main(){

	Input := []int{1,2,3,4,5,6,7}
	// fmt.Println(rotate(Input,6))
	rotate(Input,3)


}

func rotate(nums []int, k int)  {
	aSize := len(nums)
	rotationPoint := (k%aSize)
	rotationPoint = (aSize - rotationPoint - 1)

	if k == 0 || len(nums) < 2{
		return
	}

	for i := 0; i < ((rotationPoint+1)/2); i++{
		nums[i], nums[rotationPoint-i] = nums[rotationPoint-i], nums[i]
	}

	temp := nums[rotationPoint+1:]
	tSize  := len(temp)

	for i := 0; i < tSize/2; i++{
		temp[i], temp[tSize-1-i] = temp[tSize-1-i], temp[i]
	}


	for i := 0; i < aSize/2; i++{
		nums[i], nums[aSize-1-i] = nums[aSize-1-i], nums[i]
	}
}