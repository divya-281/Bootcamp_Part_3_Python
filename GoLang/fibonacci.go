package main

import "fmt"
func fibinocci(a int) int{
	if a==0{
		return 0
	}else if a==1{
		return 1
		
	}else{
		return fibinocci(a-1) +fibinocci(a-2)

	}
}
func main3() {
	for i:=0;i<6;i++{
	
		fmt.Println(fibinocci((i)))
	}

}