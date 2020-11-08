/*
	This example deals with Structs
*/

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

	aman := person{firstName: "Aman", lastName: "Jaiswal"}

	/*
		IMPORTANT:
		Passing variable like this won't update its value
		As in Go, call by value is done in such cases
		Meaning Go on calling such functions create a copy of the data provided
		Hence no change in the original value takes place.
	*/
	aman.updateName("Amit")
	fmt.Println("No change taking place on updating name")
	aman.printPerson()
	fmt.Println("----------")

	/*
		FIX:
		To pass the address of the variable instead of value
		By this the change made is permanent as it took place
		On a variable's address/reference and not on its value
	*/
	amanPointer := &aman
	amanPointer.updateNameByRef("Amit")
	fmt.Println("Change takes place on accessing via address")
	aman.printPerson()
	fmt.Println("----------")

}

func (p person) printPerson() {
	fmt.Println(p)
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

// Address pointer being passed which is accessed via *
func (p *person) updateNameByRef(newName string){
	(*p).firstName = newName
}