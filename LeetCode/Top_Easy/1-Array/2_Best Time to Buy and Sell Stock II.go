package main

import "fmt"


func main(){

	Input := []int{7,1,5,3,6,7,8,9}
	fmt.Println(maxProfit(Input))
}

func maxProfit(prices []int) int {
    var profit int
    for i := 0; i < len(prices)-1; i++ {
        if prices[i] < prices[i+1] {
            profit += prices[i+1] - prices[i]
        }
    }
    
    return profit
}