package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "time"
)

func main() {
        arguments := os.Args // providing same port as client is expecting to connect on
        // Ex -  go run tcpS.go 3000
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        // get the port from argument
        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp", PORT) // on current IP, start listening on PORT
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close() // On exiting, close the server

        c, err := l.Accept() // Accept any new connection, ex -  client's requests
        if err != nil {
                fmt.Println(err)
                return
        }

        // infinite loop
        for {
                netData, err := bufio.NewReader(c).ReadString('\n') // get data sent by client
                if err != nil {
                        fmt.Println(err)
                        return
                }
                // if it sends STOP, close server and call defer
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData)) // print data received from client
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                c.Write([]byte(myTime)) // Send the string containing time to same client
        }
}