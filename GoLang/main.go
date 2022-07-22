// main.go file
/* this
is a
multi line comment */
package main

import "fmt"
func main() {
fmt.Println("Hello Go")

// Declare variable in GoLang
 var x int =7;
 fmt.Println(x);

 //Declare variable in GoLang second method
y := 10;
fmt.Println(y);

//Declare multiple variables
var a, b, c int =1,2,3
fmt.Println(a,b,c)

// Declare variables as constant using const keyword
const pi =3.14;

//Declaration of array in GoLang
//1
var array1=[4]int{1,2,3,4}
//2
array2:=[4] int {5,6,7,8}

fmt.Println(array1)
fmt.Println(array2)

// indexing
fmt.Println(array2[2])

//mutable
array1[0]=10
fmt.Println(array1[0])

//empty array
array3 :=[10] int {}
array4 :=[10]int {10,20}

array5 :=[10]int {1,2,3,4,5,6,7,8,9,10}

fmt.Println(array3)
fmt.Println(array4)
fmt.Println(array5)

array6:=[4]int{3:30, 2:20}
fmt.Println(array6)

//length of array
fmt.Println(len(array6))

// first way of creating a slice
slice:=[] int {1,2,3,4}
fmt.Println(slice)

//length of slice
fmt.Println(len(slice))

slice2:=[] int {}

//capacity of slice
fmt.Println(cap(slice2))

fmt.Println(slice2)

slice3:=[]string{"a","b","c","d"}
fmt.Println(slice3)
fmt.Println(len(slice3))
fmt.Println(cap(slice3))

// slice from array
array7 :=[8] int {1,2,3,4,5,6,7,8}
slice4 := array7[2:4]
fmt.Println(slice4)
fmt.Println(len(slice4))
fmt.Println(cap(slice4))

slice5:=array7[0:8]
fmt.Println(slice5)
fmt.Println(len(slice5))
fmt.Println(cap(slice5))

//slice using make()
slice6 := make([] int, 4 , 8)
fmt.Println(slice6)
fmt.Println(len(slice6))
fmt.Println(cap(slice6))

//indexing of slice is same as array

//append()
slice7 :=[] string {"e","f","g","h"}
slice7 = append(slice7,"i","j","k" )
fmt.Println(slice7)
}
//    go run main.go



