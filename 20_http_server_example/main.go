package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main()  {
	// which function to revert calls to when "/" is encountered
	http.HandleFunc("/", handler)

	// where to run the server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type User struct { 
	Id    int     `json:"id"` 
	Name  string  `json:"name"` 
	Email string  `json:"email"` 
	Phone string  `json:"phone"` 
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // IMPORTANT

	user := User {
		Id: 1, 
		Name: "John Doe", 
		Email: "johndoe@gmail.com", 
		Phone: "000099999",
   	} 

	// Return as JSON Response on every request on "localhost:8000"
	json.NewEncoder(w).Encode(user) 
}