/*
	This example deals with
	1.Creating maps
	2.Printing maps
	3.Deleteing keys from map
	4.Accessing values from map via keys
	5.Different way to create a map
	6.Looping over map values
	7.
*/

package main

import "fmt"



/*
	IMPORTANT:
	All the keys and values of any map
	Are the same
	I.e. all keys should be int
	All keys should be string
	All values should be string at a time and so on
*/
func main() {
	// Extra comma at the end
	colors := map[string]string{
		"black": "#FFFFFF",
		"white": "#000000",
	}
	fmt.Println(colors)


	// Other ways to delcare empty map (to be used later to save values)
	var colors2 map[string]string
	fmt.Println(colors2)

	// Another way using make function
	colors3 := make(map[int]string)
	colors3[1] = "#FFFFFF"
	fmt.Println(colors3)
	// Doesn't support dot notation for value access
	// So all keys string/int, should be passed same
	fmt.Println(colors3[1])

	fmt.Println(colors["black"])

	/*
		Deleting a key from a map
	*/
	delete(colors, "black")
	fmt.Println(colors)


	// For loop over a map
	colors4 := map[int]string {
		1: "Red",
		2: "Green",
		3: "Blue",
	}	


	printMap(colors4)
}



// Function to be used to print map
func printMap(c map[int]string){
	for key, value := range c {
		fmt.Println(key, "has value", value)
	}
}


/*
	IMPORTANT:
	Maps vs Structs
	Map, Same type key/values vs Struct, different type values are acceptable
	Map, Reference type variable vs Struct, Value type, hence require address to change values
	Map, Used for closely related properties, Ex - Colors, vs Struct, handle more dynamic values (-ish)
*/