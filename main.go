package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn){
	buf := make([]byte, 1024)
	// We can extract the http url and write the further logic for methods and all, just like other servers like flask and all does 
	_,err := conn.Read(buf);
	if err!=nil {
		log.Fatal("Error reading: ", err);
	}

	time.Sleep(1*time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, world!\r\n"));

	conn.Close();
}

// Creating a TCP server from scratch
func main(){

	// Step 1: start listening on a port
	listener, err := net.Listen("tcp", ":5723");
	if err!=nil {
		log.Fatal("Error listening: ", err);
	}

	// Step 2: Wait for client to connect. 'Accept' is a blocking system call
	// conn, err := listener.Accept();
	// if err!=nil {
	// 	log.Fatal("Error accepting: ", err);
	// }

	// Step 3: Read the request data, Write and send the response, Close the connection
	// do(conn)

	// Step 4: Do it over and over again
	for{
		conn, err := listener.Accept();
		if err!=nil {
			log.Fatal("Error accepting: ", err);
		}
		// do(conn)

		// Step 5: multiple connections parallel
		go do(conn)
	}

}