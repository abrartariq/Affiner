package main

import("fmt")

func climbStairs(n int) int {
	solutions := make([]int, n+1)
    return climbDP(n,solutions)
}

func climbDP(n int, solution []int) int{
    if n < 0{
     return 0   
    }else if n == 1 || n == 0{
		return 1
	}else if n == 2{
		return 2
	}else if solution[n] != 0{
		return solution[n]
	}else{
		nSolution := climbDP(n-1,solution) + climbDP(n-2,solution)
		solution[n] = nSolution
	}
	return solution[n]			
}	




func main(){
	val := make([]int, 10)
	test(val)
	fmt.Println(val)
}