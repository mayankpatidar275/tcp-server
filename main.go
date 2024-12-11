package main

import (
	"fmt"
	"log"
	"net"
)

// Creating a TCP server from scratch
func main(){

	// Step 1: start listening on a port
	listener, err := net.Listen("tcp", ":5723");
	if err!=nil {
		log.Fatal("Error listening: ", err);
	}

	// Step 2: Wait for client to connect. 'Accept' is a blocking system call
	conn, err := listener.Accept();
	if err!=nil {
		log.Fatal("Error accepting: ", err);
	}

	fmt.Println(conn)

}