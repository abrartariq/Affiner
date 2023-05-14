import "fmt"


func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, k := m-1, n-1, m+n-1

	for i > -1 && j > -1 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		}else{
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
    
    if j > -1{
        for idx,val := range nums2[:j+1]{
			nums1[idx] = val
		}
    }
}


func main(){

}


