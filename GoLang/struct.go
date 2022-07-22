package main

import "fmt"

	
type Employee struct{
	fname string
	age int
	salary float64

}
func main4() {
	var emp1 Employee
	var emp2 Employee
	var emp3 Employee

	emp1.fname="Divya"
	emp1.age=22
	emp1.salary=4000

	emp2.fname="Divya1"
	emp2.age=23
	emp2.salary=2000

	emp3.fname="Divya2"
	emp3.age=25
	emp3.salary=3000


	fmt.Println("Name of Employee :", emp1.fname, "Age :",emp1.age, "Salary :",emp1.salary)
	fmt.Println("Name of Employee :", emp2.fname, "Age :",emp2.age, "Salary :",emp2.salary)
	fmt.Println("Name of Employee :", emp3.fname, "Age :",emp3.age, "Salary :",emp3.salary)
}

// go run struct.go