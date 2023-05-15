import "fmt"


/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	low := 1
	high := n
	mid := 1 + (high-1)/2
	for low < high{
		if isBadVersion(mid) {
			high = mid
			mid = low + (high-low)/2
		}else{
			low = mid + 1
			mid = low + (high-low)/2
		}
	}
	return mid
}




func main(){

}


