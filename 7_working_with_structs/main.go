package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email string
	zipCode int
}

type personWithContact struct {
	firstName string
	lastName string
	contact contactInfo
}


func main() {

	// Initializng a struct with values
	aman := person{firstName: "Aman", lastName: "Jaiswal"}
	fmt.Println(aman)
	fmt.Println("-------")
	fmt.Println("Printing Empty struct")

	var aman2 person
	/*
		When initialized with no values, a struct variable fills its properties
		With null values, i.e. "" for Strings, 0 for Int/Float and False for Bool
	*/
	fmt.Println(aman2)
	fmt.Println("-------")

	// Assigning values to a struct variable
	fmt.Println("Printing Struct with assigned values")
	aman2.firstName = "Aman"
	aman2.lastName = "Jaiswal"
	fmt.Println(aman2)

	// Using struct as type in another struct
	fmt.Println("Using struct as datatype in another")

	/*
		In initialization as done below,
		Its compulsory to put , after every property assignment
		Even after the last one in every struct declaration
	*/
	aman3 := personWithContact{
		firstName: "Aman",
		lastName: "Jaiswal",
		contact: contactInfo {
			zipCode: 2080,
			email: "aman@gmail.com",
		},
	}
	fmt.Println(aman3)
	fmt.Println("-------")
}
