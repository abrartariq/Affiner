package main

import "fmt"

func main() {

	tempInput := []int{1, 2, 3, 4}
	fmt.Println(removeDuplicates(tempInput))
}

func removeDuplicates(nums []int) int {
	a := 0

	for b := 1; b < len(nums); b++ {
		if nums[a] != nums[b] {
			a++
			nums[a] = nums[b]
		}
	}

	return a + 1
}
