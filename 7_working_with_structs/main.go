/*
	This example deals with Structs
	1. Creating structs
	2. Adding methods for structs passing struct as receiver
	3. Using one struct as datatype for another
	Creations: person, personWithContact, printPerson()
*/

package main

import "fmt"

// Custom type struct
type person struct {
	firstName string
	lastName  string
}

type contactInfo struct {
	email string
	zipCode int
}

// Using contactInfo as datatype for this
type personWithContact struct {
	firstName string
	lastName string
	// Alternative: 'contactInfo', 
	// hence variable and its datatype are both contactInfo
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
	fmt.Println("-------")
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
	fmt.Println("Using custom methods with structs, Ex- print a person struct")
	aman.printPerson()


}

// Since this accepts person as a reciever, it will be available after
// All person type structs
func (p person) printPerson() {
	fmt.Println(p)
}