package main

import (
    "fmt"
    "bufio"
    "os"
    "net"
)

const (
    defaultHost = "localhost"
    defaultPort = 9999
)

// To test your server implementation, you might find it helpful to implement a
// simple 'client runner' program. The program could be very simple, as long as
// it is able to connect with and send messages to your server and is able to
// read and print out the server's response to standard output. Whether or
// not you add any code to this file will not affect your grade.
func main() {

    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", defaultHost, defaultPort))
    if err != nil {
        fmt.Printf("%v", err)
    }

    for {
        // read in input from stdin
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        // send to socket
        fmt.Fprintf(conn, text + "\n")
        // listen for reply
        //message, _ := bufio.NewReader(conn).ReadString('\n')
        //fmt.Print("Message from server: "+message)
    }

}
