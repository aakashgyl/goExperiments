package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//We will use this wire protocol: (this is open connection unless close)
//hello -> hi
//time -> 10:48am
//close

func main() {
	srv, err := net.Listen("tcp", ":8087") //gives a server object which is not listening
	defer srv.Close()
	if err != nil {
		log.Printf("Cannot start server %s", err)
		return
	}

	log.Println("Server started")
	for {
		conn, err := srv.Accept() //this will make it listen. Its a blocking call and waits for a request
		if err != nil {
			log.Printf("Error: %v\n", err)
			return
		}

		go handleConnection(conn)
	}
	log.Println("Server ended")
}

func handleConnection(conn net.Conn) {
	log.Println("Connection started...")
	var command string

	outerloop:
	for {
		log.Println("Let me wait")
		_, err := fmt.Fscanln(conn, &command)
		log.Println("Let me start")

		if err != nil {
			log.Printf("Conn err: %v\n", err)
			_, ok := err.(net.Error)
			if ok {
				break outerloop
			}

			if err == io.EOF {
				break outerloop
			}
		}

		switch command {
		case "hello":
			fmt.Fprintln(conn, "Hi")
		case "time":
			fmt.Fprintln(conn, time.Now())
		case "close":
			conn.Close()
			break outerloop
		default:
			fmt.Println("Invalid use of TP protocol")
		}
	}
	log.Println("Connection closed...")
}
