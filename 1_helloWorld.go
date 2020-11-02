// To create an executable package, we always put package as name "main"
// And a mandatory function named main()
// To create a helper type package, we will use a different name for package

// NOTE: 'go build' only creates a package when package is main, and not for anything else.
package main

// To import a package (reusable??), import is used
// fmt or format is a standard library package, containing several methods to print data
import "fmt"

// To create a function
func main(){
	fmt.Println("Hello World!")
}

// go run <FILE_NAME> -- compiles and executes a file
// go build -- creates an executable under folder_name that can be run by ./<FOLDER_NAME>
