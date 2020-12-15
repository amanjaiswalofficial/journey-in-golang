package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	// Connect with a IP i.e. host:port, as given as cli arg
	// Ex - go run tcpC.go 127.0.0.1:3000
	if len(arguments) == 1 {
			fmt.Println("Please provide host:port.")
			return
	}

	CONNECT := arguments[1] // get the host:port
	// call to TCP Server running on host:port as given in CONNECT
	c, err := net.Dial("tcp", CONNECT) 
	if err != nil {
			fmt.Println(err)
			return
	}

	// infinite loop
	for {
			
			reader := bufio.NewReader(os.Stdin) // start a shell reader to read input
			fmt.Print(">> ") // print to receive input
			text, _ := reader.ReadString('\n') // read the string entered in shell


			fmt.Fprintf(c, text+"\n") // Send the string to the connected server (i.e. c)

			message, _ := bufio.NewReader(c).ReadString('\n') // get the response from server
			fmt.Print("->: " + message) // print the same

			// if command entered is STOP, then close the 
			if strings.TrimSpace(string(text)) == "STOP" {
					fmt.Println("TCP client exiting...")
					return
			}
	}
}