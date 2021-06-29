package main

import (
	"fmt"
	"log"
	"net"
)

func handle_conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connect ",conn.RemoteAddr())
	//communication
	buf := make([]byte,256)
	for {
		//read from network
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read",err)
			break
		}
		fmt.Println(n, string(buf[:n]))
		//write back
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Failed to Write", err)
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp","localhost:8888")
	if err != nil {
		log.Panic("Failed to Listen", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept ", err)
		}
		go handle_conn(conn) //connect with new connection
	}
}
