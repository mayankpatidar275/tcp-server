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

	fmt.Println("Hello world!", listener)

}