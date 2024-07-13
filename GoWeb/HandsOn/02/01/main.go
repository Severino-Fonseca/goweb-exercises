package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panicln("Shit")
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("error accepting connection", err)
		}

		io.WriteString(conn, "i see you conected\n")

		conn.Close()
	}
}
