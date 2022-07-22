package main

import "fmt"
func factorial(a int) int{
if a==1{
		return 1
		
	}else{
		return a*factorial(a-1)

	}
}
func main1() {
	
	
		fmt.Println(factorial((5)))
	}

