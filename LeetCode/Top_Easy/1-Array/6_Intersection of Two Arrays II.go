package main

import "fmt"


func main(){

	Input := []int{1,1,3,5,6,3,5}
	InputB := []int{1,1,3,5,6,3,5}
	fmt.Println(intersect(Input,InputB[:5]))


}


func intersect(nums1 []int, nums2 []int) []int {
 
    var result []int

    size1 := len(nums1)
    size2 := len(nums2)

	quicksort(nums1)
	quicksort(nums2)

    smallArr := nums1
    bigArr := nums2

    if size1 > size2{

    	smallArr = nums2 
    	bigArr = nums1
    	size1, size2 = size2, size1
    }

    for j, i := 0, 0; i < size1 && j <size2 ;  {
    	if smallArr[i] == bigArr[j]{
    		result = append(result,smallArr[i])
    		i++
    		j++
    	}else if smallArr[i] < bigArr[j] {
    		i++
    	}else{
    		j++
    	}
    }
		    	
    return result

}


func median(valToSort []int, a int, b int, c int) int{

	if valToSort[b] < valToSort[c] {
		// b < c
		if valToSort[a] < valToSort[b] {
			// a < b < c
			return b
		}else if valToSort[a] < valToSort[c] {
			// b < a < c
			return a
		}else{
			// b < c < a
			return c
		}
	}else{
		// c < b
		if valToSort[a] < valToSort[c] {
			// a < c < b
			return c
		}else if valToSort[a] < valToSort[b] {
			// c < a < b
			return a
		}else{
			// c < b < a
			return b
		}
	}
}

func quickSortHelper(valToSort []int, left int, right int){

	if left >= right {
		return
	}

	var temp int = -99
	i := left
	j := right
	mid := int(left + (right-left)/2)
	pivot := valToSort[median(valToSort, left, mid, right)]

	for i <= j {

	    for valToSort[i] < pivot {i++;}
	    for valToSort[j] > pivot {j--;}

	    if i <= j {

	          temp = valToSort[i]
	          valToSort[i] = valToSort[j]
	          valToSort[j] = temp
	          i++
	          j--

	    }

	}

	//Recursion//
	quickSortHelper(valToSort, left, j)
	quickSortHelper(valToSort, i, right)

}


func quicksort(valToSort []int) []int{

	size := len(valToSort)

	result := make([]int, size, size)


	quickSortHelper(valToSort, 0, size-1)

	for i := 0; i < size; i++ {
		result[i] = valToSort[i]
	}


	return result
}

