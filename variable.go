package main

import "fmt"

func main() {
	//variable declearation
	var fullName, address string // we can decleare similar type of variable by separating them with comma
	var age int

	var applePrice int = 5 // We can decleare and initialize variable in same line also

	var salary = 10.25 // Even we need not to assign the data type, Because Go can get it out

	fruitsName := "Orange" // Moreover we can also exclude "var" by putting ":" before assigning value after "="

	//variable initialization
	fullName = "Yksi Pilvista"
	age = 32
	address = "Hervanta"

	fmt.Println("Name of Employee: ", fullName)
	fmt.Println("Age: ", age)
	fmt.Println("Living Address: ", address)
	fmt.Println("Hourly Salary: ", salary)
	fmt.Println("Apple Price:", applePrice)
	// We can format a line as below also by using "Printf" and "%v" for variable:
	fmt.Printf("Favourite Fruit:%v, But Apple price of today: %v\n", fruitsName, applePrice)
	fmt.Print("Keep good communication")
}
