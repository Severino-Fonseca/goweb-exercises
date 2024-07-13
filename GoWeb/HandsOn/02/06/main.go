package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	var m, uri string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			m = xs[0]
			uri = xs[1]
			fmt.Println(m)
			fmt.Println(uri)
		}
		if ln == "" {
			break
		}
		i++
	}

	body := "answers in payload\n"
	body += m
	body += "\n"
	body += uri

	io.WriteString(conn, "HTTP1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
