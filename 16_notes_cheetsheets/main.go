package main

import "fmt"

func main() {
	// Formatting examples
	fmt.Printf("%T %v\n", 10, 10)
	fmt.Printf("%v\n", false)

	fmt.Printf("%e\n", 1.23456789999)
	fmt.Printf("%f\n", 1.234567)

	fmt.Printf("%s\n", "ABC")
	fmt.Printf("%q\n", "ABC with quotes")

	fmt.Printf("%9q\n", "Aman")
	fmt.Printf("%-9q", "Aman")
	fmt.Printf("%q\n", "Aman")
}
