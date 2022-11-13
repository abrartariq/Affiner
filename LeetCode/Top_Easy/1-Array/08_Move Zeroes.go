package main

import "fmt"

func main(){
	Input := []int{0,0,0,0,0,0,1,1,1}
	// fmt.Println(moveZeroes(Input))
	moveZeroes(Input)
	fmt.Println(Input)
}

// Efficient
func moveZeros(nums []int){

	fstIter := 0
	aSize := len(nums)

	for i := 0; i < aSize; i++ {
		if nums[i] != 0{
			nums[fstIter] = nums[i]
			fstIter++
		}
	}

	for i := fstIter; i < aSize; i++ {
		nums[i] = 0
	}


}

// First 
func moveZeroes(nums []int)  {
	sizeA := len(nums)

	iZero, inZero := findNextnz(nums)

	if sizeA < 2 || iZero > (sizeA-2){
		return
	}

	if nums[0] == 0 {
		iZero, inZero = inZero, iZero
		nums[iZero], nums[inZero] = nums[inZero], nums[iZero]
	}


	for i := inZero; i < sizeA-1; i++ {
		if nums[i+1] != 0{
			inZero++
		}else{
			break
		}
	}

	for i := 1; i < sizeA; i++ {
		if nums[i] != 0{
			continue
		}else{

			inZero++;
			j := i 
			for ; j < sizeA; j++ {
				if nums[j] ==  0{
					continue
				}else{
					break
				}
			}
			if j == sizeA{
				break
			}else{
				nums[inZero], nums[j] = nums[j], nums[inZero]
			}
		}
	}
}

func findNextnz(nums []int) (int, int){

	sizeA := len(nums)
	iZero, inZero := 0, 0
	flagZ, flagN := true, true

	for i := 0; i < sizeA && (flagZ || flagN); i++ {
		if flagZ && nums[i] == 0 {
			iZero = i
			flagZ = false
		}else if flagN && nums[i] != 0 {
			inZero = i
			flagN = false
		}
	}

	return iZero, inZero
}
