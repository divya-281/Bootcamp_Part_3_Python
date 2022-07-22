package main

import "fmt"

func function1(){
	fmt.Println("Hello I'm in function1")
}
func function2(name string , age int){
	fmt.Println("Hello I am", name,"and I'm", age)
}
func sum(a int, b int ) int{
	return a+b
}


func main2() {

//larger of two values
a:= 76
b:=90
if a>b{
	fmt.Println("a is greater")
} else 
{
	fmt.Println("b is greater")
}

//switch case
day :=6
switch(day){
case 1:
fmt.Println("Sunday")
case 2:
	fmt.Println("Monday")
case 3:
	fmt.Println("Tuesday")
case 4:
	fmt.Println("Wednesday")
case 5:
	fmt.Println("Thursday")
case 6:
	fmt.Println("Friday")
case 7:
	fmt.Println("Saturday")
default:
	fmt.Println("Invalid Input")
}

//for loop syntax, break, continue
for i:=0;i<11;i++{
	if i==3{
		continue
	}
	if i==8{
		break
	}
	fmt.Println(i)
}

//function creation in GoLang
//define outside main method
function1()
function2("Divya", 22)

// function with return type
fmt.Println(sum(80,1800))




}
//  go run main2.go